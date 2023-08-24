package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
	"unicode"

	yaml "gopkg.in/yaml.v3"
)

const dockerDocsZip = "https://github.com/docker/docs/archive/refs/heads/main.zip"

type ctrCliCmd struct {
	Command          string
	Short            string
	Long             string
	Usage            string
	Pname            string
	Plink            string
	Deprecated       bool
	MinApiVersion    string
	Experimental     bool
	Experimentalcli  bool
	Kubernetes       bool
	Swarm            bool
	Subcommand       []string
	Options          []*ctrCliOpt
	InheritedOptions []*ctrCliOpt `yaml:"inherited_options"`
	Arguments        []*ctrCliArg
}

type ctrCliArg struct {
	Name     string
	Optional bool
	Multiple bool
}

type ctrCliOpt struct {
	Option          string
	ValueType       string `yaml:"value_type"`
	Description     string
	Deprecated      bool
	Hidden          bool
	Experimental    bool
	Experimentalcli bool
	Kubernetes      bool
	Swarm           bool
	DefaultValue    string `yaml:"default_value"`
}

var cmdOverrides = map[string]*ctrCliCmd{
	"cp": {
		Arguments: []*ctrCliArg{
			{Name: "srcpath"},
			{Name: "dstpath"},
		},
	},
	"container cp": {
		Arguments: []*ctrCliArg{
			{Name: "srcpath"},
			{Name: "dstpath"},
		},
	},
}

var cmdTemplate = `package ctrctl

type {{ .FuncName }}Opts struct {
{{ .OptsStruct }}
}

// {{ .Short }}
func {{ .FuncName }}(opts *{{ .FuncName }}Opts{{if .ArgsDefn}}, {{ .ArgsDefn }}{{end}}) (
	stdout string, stderr string, err error) {
	return runCtrCmd(
		{{ .Subcommand }},
		{{if .Args}}{{ .Args }},
{{end}}		opts,
		{{ .OptsPos }},
	)
}
`

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	docsDir, err := os.MkdirTemp("", "ctrctl.*")
	if err != nil {
		return fmt.Errorf("error creating temporary directory: %w", err)
	}
	defer os.RemoveAll(docsDir)

	fetchZip(dockerDocsZip, docsDir)

	tmpl, err := template.New("").Parse(cmdTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	err = filepath.WalkDir(
		filepath.Join(docsDir, "docs-main", "data", "engine-cli"),
		func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			buf, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file: %w", err)
			}
			addCliDefinition(buf, tmpl)
			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("error gathering cli data: %w", err)
	}

	return nil
}

func fetchZip(url string, dir string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading downloaded file: %w", err)
	}

	archive, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return fmt.Errorf("error reading zip: %w", err)
	}

	err = extractZip(archive, dir)
	if err != nil {
		return fmt.Errorf("error extracting zip: %w", err)
	}

	return nil
}

func extractZip(archive *zip.Reader, dir string) error {
	for _, f := range archive.File {
		fpath := filepath.Join(dir, f.Name)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, 0755)
			if err != nil {
				return fmt.Errorf("error creating dir (%s): %w", fpath, err)
			}
			continue
		}

		file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		freader, err := f.Open()
		if err != nil {
			return fmt.Errorf("error opening file (%s): %w", f.Name, err)
		}

		_, err = io.Copy(file, freader)
		if err != nil {
			return err
		}

		file.Close()
		freader.Close()
	}
	return nil
}

func addCliDefinition(buf []byte, tmpl *template.Template) error {
	cmd, err := cmdFromDefinition(buf)
	if err != nil {
		return err
	}

	trimmedCmd := strings.TrimPrefix(cmd.Command, "docker ")
	if trimmedCmd == "" {
		return nil
	}

	filename := underscore(trimmedCmd) + ".go"
	out, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error creating %s: %w", filename, err)
	}
	defer out.Close()

	return tmpl.Execute(out, struct {
		FuncName   string
		OptsStruct string
		Short      string
		ArgsDefn   string
		Args       string
		OptsPos    int
		Subcommand string
	}{
		FuncName:   cmd.methodName(),
		OptsStruct: cmd.optsStruct(),
		Short:      ensureDot(strings.TrimSpace(cmd.Short)),
		ArgsDefn:   cmd.argsDefn(),
		Args:       cmd.argsSlice(),
		OptsPos:    cmd.optsPos(),
		Subcommand: cmd.subcommandSlice(),
	})
}

func cmdFromDefinition(buf []byte) (*ctrCliCmd, error) {
	cmd := &ctrCliCmd{}
	err := yaml.Unmarshal(buf, &cmd)
	if err != nil {
		return nil, err
	}

	usage := strings.TrimPrefix(cmd.Usage, "docker ") + " "
	var word []rune
	var depth int
	var dots int
	var subFinished bool
	var nextUpper bool
	arg := &ctrCliArg{}
	for _, r := range usage {
		if r != '.' {
			dots = 0
		}
		switch r {
		case ' ':
			if depth == 0 {
				if len(word) == 0 {
					nextUpper = false
					continue
				}
				if subFinished {
					arg.Name = string(word)
					cmd.Arguments = append(cmd.Arguments, arg)
					arg = &ctrCliArg{}
				} else {
					cmd.Subcommand = append(cmd.Subcommand, string(word))
				}
				word = []rune{}
			} else {
				nextUpper = true
			}
		case '[':
			subFinished = true
			if depth == 0 {
				arg.Optional = true
			}
			depth++
		case ']':
			depth--
		case '.':
			dots++
			if dots > 2 {
				arg.Multiple = true
			}
		default:
			if subFinished && !unicode.IsLetter(r) && !unicode.IsNumber(r) {
				nextUpper = true
				continue
			}
			if unicode.IsUpper(r) {
				subFinished = true
			}
			if nextUpper {
				nextUpper = false
				r = unicode.ToUpper(r)
			} else {
				r = unicode.ToLower(r)
			}
			word = append(word, r)
		}
	}

	cmd.deduplicateArgs()
	cmd.convertOptNames()

	overrides, ok := cmdOverrides[strings.TrimPrefix(cmd.Command, "docker ")]
	if ok {
		overrideStruct(cmd, overrides)
	}

	return cmd, nil
}

func hasLowercase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func (cmd *ctrCliCmd) deduplicateArgs() {
	names := make(map[string]bool)
	for _, arg := range cmd.Arguments {
		_, seen := names[arg.Name]
		if seen {
			arg.Name = fmt.Sprintf("extra%s", capitalize(arg.Name))
		}
		names[arg.Name] = true
	}
}

func (cmd *ctrCliCmd) convertOptNames() {
	for i := 0; i < len(cmd.Options); i++ {
		var name []rune
		var nextUpper bool
		for _, r := range cmd.Options[i].Option {
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				if nextUpper {
					nextUpper = false
					r = unicode.ToUpper(r)
				}
				name = append(name, r)
			} else {
				nextUpper = true
			}
		}
		cmd.Options[i].Option = string(name)
	}
}

func (cmd *ctrCliCmd) methodName() string {
	parts := strings.Fields(strings.TrimPrefix(cmd.Command, "docker "))
	var result []string

	for _, part := range parts {
		titled := []rune{}
		nextUpper := true
		for _, r := range part {
			if nextUpper {
				r = unicode.ToUpper(r)
				nextUpper = false
			}
			if !unicode.IsLetter(r) {
				nextUpper = true
				continue
			}
			titled = append(titled, r)
		}
		result = append(result, string(titled))
	}

	return strings.Join(result, "")
}

func (cmd *ctrCliCmd) optsStruct() string {
	var s strings.Builder
	for _, opt := range cmd.Options {
		for _, line := range strings.Split(opt.Description, "\n") {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			s.WriteString(fmt.Sprintf("\t// %s\n", ensureDot(line)))
		}
		s.WriteString(fmt.Sprintf("\t%s string\n\n", capitalize(opt.Option)))
	}
	return strings.TrimRight(s.String(), "\n")
}

func (c *ctrCliCmd) optsPos() int {
	for i, arg := range c.Arguments {
		if arg.isOpts() {
			return i
		}
	}
	return -1
}

func (c *ctrCliCmd) argsDefn() string {
	var defns []string
	for _, arg := range c.Arguments {
		if arg.isOpts() {
			continue
		}
		var defn string
		if arg.Multiple {
			defn = fmt.Sprintf("%s ...string", arg.Name)
		} else {
			defn = fmt.Sprintf("%s string", arg.Name)
		}
		defns = append(defns, defn)
	}
	return strings.Join(defns, ", ")
}

func (c *ctrCliCmd) argsSlice() string {
	var defns []string
	var vararg bool
	for _, arg := range c.Arguments {
		if arg.isOpts() {
			continue
		}
		if arg.Multiple {
			vararg = true
			continue
		}
		defns = append(defns, arg.Name)
	}
	if len(defns) == 0 {
		return "[]string{}"
	} else if vararg {
		return fmt.Sprintf("append([]string{ %s }, %s...)",
			strings.Join(defns, ","),
			c.Arguments[len(c.Arguments)-1].Name,
		)
	} else {
		return fmt.Sprintf("[]string{ %s }", strings.Join(defns, ", "))
	}
}

func (c *ctrCliCmd) subcommandSlice() string {
	quoted := []string{}
	for _, part := range c.Subcommand {
		quoted = append(quoted, "\""+part+"\"")
	}
	return fmt.Sprintf("[]string{ %s }", strings.Join(quoted, ", "))
}

func (a *ctrCliArg) isOpts() bool {
	return a.Name == "options"
}

func capitalize(s string) string {
	var result []rune
	for i, r := range s {
		if i == 0 {
			r = unicode.ToUpper(r)
		}
		result = append(result, r)
	}
	return string(result)
}

func underscore(s string) string {
	var result []rune
	var under bool
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			result = append(result, r)
			under = false
		} else if !under {
			result = append(result, '_')
			under = true
		}
	}
	return string(result)
}

func ensureDot(s string) string {
	if strings.HasSuffix(s, ".") {
		return s
	}
	return s + "."
}

func overrideStruct(src, dst interface{}) {
	srcv := reflect.ValueOf(src).Elem()
	dstv := reflect.ValueOf(dst).Elem()

	for i := 0; i < dstv.NumField(); i++ {
		srcf := srcv.Field(i)
		dstf := dstv.Field(i)

		if !dstf.IsZero() {
			srcf.Set(dstf)
		}
	}
}
