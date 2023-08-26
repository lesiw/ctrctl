package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
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
	Varg             *ctrCliArg
}

type ctrCliArg struct {
	Name         string
	Optional     bool
	Varg         bool
	VargRequired bool
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

var cmdTemplate = `type {{ .FuncName }}Opts struct {
	// Base exec.Cmd.
	Cmd *exec.Cmd

{{ .OptsStruct }}
}

// {{ .Short }}
func {{ .FuncName }}(opts *{{ .FuncName }}Opts{{if .ArgsDefn}}, {{ .ArgsDefn }}{{end}}) (string, error) {
{{if .VargRequired }}	if len({{ .Varg }}) == 0 {
		return "", fmt.Errorf("{{ .Varg }} must have at least one value")
	}
{{end}}	return runCtrCmd(
		{{ .Subcommand }},
{{if .Args}}		{{ .Args }},
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

	if err := clearGeneratedFiles(); err != nil {
		return fmt.Errorf("error clearing previously generated files: %w", err)
	}
	if err := fetchZip(dockerDocsZip, docsDir); err != nil {
		return fmt.Errorf("error fetching docker docs: %w", err)
	}

	tmpl, err := template.New("").Parse(cmdTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	cliDocsDir := filepath.Join(docsDir, "docs-main", "data", "engine-cli")
	m, err := outfilemap(cliDocsDir)
	if err != nil {
		return fmt.Errorf("error mapping output files: %w", err)
	}
	for name, files := range m {
		out, err := os.OpenFile(name+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", name+".go", err)
		}
		defer out.Close()
		if _, err := out.WriteString("package ctrctl\n\n"); err != nil {
			return fmt.Errorf("error writing file %s: %w", name+".go", err)
		}
		for _, file := range files {
			path := filepath.Join(cliDocsDir, file)
			in, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", path, err)
			}
			if err := addCliDefinition(in, out, tmpl); err != nil {
				return fmt.Errorf("error adding cli definition: %w", err)
			}
		}
	}

	cmd := exec.Command("gofmt", "-w", "-s", ".")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running gofmt: %w", err)
	}
	cmd = exec.Command("goimports", "-w", ".")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running goimports: %w", err)
	}

	return nil
}

func clearGeneratedFiles() error {
	files, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("error reading files in directory: %w", err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if !strings.HasSuffix(f.Name(), ".go") {
			continue
		}
		if f.Name() == "main.go" {
			continue
		}
		err := os.Remove(f.Name())
		if err != nil {
			return fmt.Errorf("failed to remove file: %w", err)
		}
	}
	return nil
}

func outfilemap(path string) (map[string][]string, error) {
	result := make(map[string][]string)
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %w", path, err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		nameparts := strings.Split(f.Name(), "_")
		var key string
		if len(nameparts) < 3 {
			key = "ctrctl"
		} else {
			key = nameparts[1]
		}
		_, ok := result[key]
		if !ok {
			result[key] = []string{}
		}
		result[key] = append(result[key], f.Name())
	}
	for name := range result {
		sort.Strings(result[name])
	}
	return result, nil
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

func addCliDefinition(in []byte, out io.Writer, tmpl *template.Template) error {
	cmd, err := cmdFromDefinition(in)
	if err != nil {
		return err
	}

	trimmedCmd := strings.TrimPrefix(cmd.Command, "docker ")
	if trimmedCmd == "" {
		return nil
	}

	var varg string
	if cmd.Varg != nil {
		varg = cmd.Varg.Name
	}

	return tmpl.Execute(out, struct {
		FuncName     string
		OptsStruct   string
		Short        string
		ArgsDefn     string
		Args         string
		OptsPos      int
		Subcommand   string
		Varg         string
		VargRequired bool
	}{
		FuncName:     cmd.methodName(),
		OptsStruct:   cmd.optsStruct(),
		Short:        ensureDot(strings.TrimSpace(cmd.Short)),
		ArgsDefn:     cmd.argsDefn(),
		Args:         cmd.argsSlice(),
		OptsPos:      cmd.optsPos(),
		Subcommand:   cmd.subcommandSlice(),
		Varg:         varg,
		VargRequired: (cmd.Varg != nil && cmd.Varg.VargRequired),
	})
}

func cmdFromDefinition(buf []byte) (*ctrCliCmd, error) {
	cmd := &ctrCliCmd{}
	err := yaml.Unmarshal(buf, &cmd)
	if err != nil {
		return nil, err
	}

	cmd.Options = append(cmd.Options, cmd.InheritedOptions...)
	cmd.convertOptNames()
	sort.Slice(cmd.Options, func(i, j int) bool {
		return cmd.Options[i].Option < cmd.Options[j].Option
	})

	cmd.parseUsage(strings.TrimPrefix(cmd.Usage, "docker ") + " ")

	if err := cmd.deduplicateArgs(); err != nil {
		return nil, fmt.Errorf("error deduplicating '%s' args: %w", cmd.Command, err)
	}

	overrides, ok := cmdOverrides[strings.TrimPrefix(cmd.Command, "docker ")]
	if ok {
		overrideStruct(cmd, overrides)
	}

	return cmd, nil
}

func (cmd *ctrCliCmd) parseUsage(usage string) {
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
				arg.Varg = true
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
}

func (cmd *ctrCliCmd) deduplicateArgs() error {
	names := make(map[string]bool)
	var lastarg *ctrCliArg
	var newargs []*ctrCliArg
	for _, arg := range cmd.Arguments {
		_, seen := names[arg.Name]
		if seen {
			if cmd.Varg != nil {
				return fmt.Errorf("multiple vargs: %s, %s", cmd.Varg.Name, arg.Name)
			}
			if lastarg != nil && lastarg.Name == arg.Name {
				cmd.Varg = lastarg
				lastarg.Varg = true
				lastarg.VargRequired = true
				continue
			} else {
				arg.Name = fmt.Sprintf("extra%s", capitalize(arg.Name))
			}
		}
		if arg.Varg {
			cmd.Varg = arg
		}
		names[arg.Name] = true
		newargs = append(newargs, arg)
		lastarg = arg
	}
	cmd.Arguments = newargs
	return nil
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
		s.WriteString(fmt.Sprintf("\t%s %s\n\n", capitalize(opt.Option), opt.structType()))
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
		if arg.Varg {
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
	for _, arg := range c.Arguments {
		if arg.isOpts() {
			continue
		}
		if arg.Varg {
			continue
		}
		defns = append(defns, arg.Name)
	}
	if len(defns) == 0 && c.Varg != nil {
		return c.Varg.Name
	} else if len(defns) == 0 {
		return "[]string{}"
	} else if c.Varg != nil {
		return fmt.Sprintf("append([]string{%s}, %s...)",
			strings.Join(defns, ","),
			c.Varg.Name,
		)
	} else {
		return fmt.Sprintf("[]string{%s}", strings.Join(defns, ", "))
	}
}

func (c *ctrCliCmd) subcommandSlice() string {
	quoted := []string{}
	for _, part := range c.Subcommand {
		quoted = append(quoted, "\""+part+"\"")
	}
	return fmt.Sprintf("[]string{%s}", strings.Join(quoted, ", "))
}

func (a *ctrCliArg) isOpts() bool {
	return a.Name == "options"
}

func (o *ctrCliOpt) structType() string {
	switch o.ValueType {
	case "bool":
		return "bool"
	case "int":
		return "*int"
	default:
		return "string"
	}
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
