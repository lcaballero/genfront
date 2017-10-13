package process

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"

	"path/filepath"

	"github.com/lcaballero/genfront/cli"
	"github.com/lcaballero/genfront/maybe"
	"github.com/spf13/viper"
)

var EnvVars = []string{
	"GOARCH",
	"GOOS",
	"GOFILE",
	"GOLINE",
	"GOPACKAGE",
	"DOLLAR",
}

type Env struct {
	pairs map[string]interface{}
}

func NewEnv() *Env {
	env := &Env{
		pairs: make(map[string]interface{}),
	}
	env.AddGoEnvironment()
	return env
}
func (e *Env) ToMap() map[string]interface{} {
	return e.pairs
}
func (e *Env) Add(key string, val interface{}) *Env {
	e.pairs[key] = val
	return e
}
func (e *Env) AddPairs(pairs map[string]interface{}) *Env {
	for k, v := range pairs {
		e.Add(k, v)
	}
	return e
}
func (env *Env) AddSettings(frontmatter string) (*Env, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	err := v.ReadConfig(bytes.NewBufferString(frontmatter))
	env.AddPairs(v.AllSettings())
	return env, err
}
func (e *Env) Sep() string {
	return strings.Repeat("-", 80)
}

func (env *Env) ShowEnvironment(w io.Writer, errs ...error) {
	for k, v := range env.pairs {
		fmt.Fprintf(w, "%s : %s\n", k, v)
	}
	for i, e := range os.Args {
		fmt.Fprintf(w, "os.Args[%d] = %s\n", i, e)
	}

	if errs != nil && len(errs) > 0 {
		for _,err := range errs {
			fmt.Fprintln(w, err)
		}
	}
}

func (env *Env) String(key string) string {
	g, ok := env.pairs[key]
	if !ok {
		return ""
	}
	s, ok := g.(string)
	if ok {
		return s
	} else {
		return ""
	}
}

// Codefile finds the .go file based on the GOFILE environment variable, else
// it uses the provided parameter: defaultGoFile.
func (env *Env) Codefile(defaultGoFile string) string {
	cwd := env.String("CWD")
	gofile := env.String("GOFILE")
	if gofile == "" {
		gofile = defaultGoFile
	}
	return filepath.Join(cwd, gofile)
}

func (env *Env) RelativeFile(f string) string {
	cwd := env.String("CWD")
	return filepath.Join(cwd, f)
}

func (env *Env) Exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func (env *Env) AddGoEnvironment() *Env {
	for _, e := range EnvVars {
		env.Add(e, os.Getenv(e))
	}

	if p, ok := env.pairs["GOPACKAGE"]; !ok || p == "" {
		env.Add("GOPACKAGE", "main")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	env.Add("CWD", wd)
	env.Add("GEN_TAGLINE", fmt.Sprintf("// Generated by %s -- do not edit this file.", os.Args[0]))
	env.Add("Args", os.Args)

	return env
}

func (env *Env) BuildFuncMap() template.FuncMap {
	return template.FuncMap{
		"pascal":    ToPascal,
		"title":     strings.Title,
		"lower":     strings.ToLower,
		"upper":     strings.ToUpper,
		"toSymbol":  ToSymbol,
		"getenv":    os.Getenv,
		"split":     strings.Split,
		"camelCase": ToCamelCase,
		"hasPrefix": strings.HasPrefix,
	}
}

func (env *Env) CreateTemplate(tpl string) (*template.Template, error) {
	return template.New("FrontMatterProcessor").Funcs(env.BuildFuncMap()).Parse(tpl)
}

func (p *Env) ShowDebug(conf *cli.CliConf, tpl *template.Template, content string) {
	w := os.Stdout
	fmt.Fprintf(w, "%s\n", p.Sep())
	p.ShowEnvironment(w)
	fmt.Fprintf(w, "%s\n", p.Sep())
	fmt.Fprintln(w, conf)
	fmt.Fprintf(w, "%s\n", p.Sep())

	if tpl != nil {
		tpl.Execute(w, p.ToMap())
	}
	if content != "" {
		fmt.Fprintf(w, content)
	}

	fmt.Fprintln(w)
}

func (p *Env) MaybeExit(conf *cli.CliConf, tpl *template.Template, content string) {
	if conf.Debug() {
		p.ShowDebug(conf, tpl, content)
	}
	if conf.Noop() {
		log.Printf("Skipping writing output file: %s", maybe.JoinCwd(conf.OutputFile()))
		os.Exit(1)
	}
}
