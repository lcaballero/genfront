package doctable

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"

	cmd "github.com/codegangsta/cli"
	"github.com/lcaballero/genfront/cli"
	"github.com/lcaballero/genfront/process"
	. "github.com/lcaballero/genfront/maybe"
	"os"
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
)


type DocFinder struct {
	*cli.CliConf
	*process.Env
}

func RunDocFinder(c *cmd.Context) {
	df := &DocFinder{
		CliConf: cli.NewCliConf(c),
		Env:     process.NewEnv(),
	}
	defer func() {
		err := recover()
		if err != nil {
			df.Env.ShowEnvironment(os.Stdout)
			fmt.Println(err)
		}
	}()
	err := df.Run()
	if err != nil {
		df.Env.ShowEnvironment(os.Stderr)
	}
}

func (d *DocFinder) findFieldDocumentation() (*FieldAndDoc, error) {
	env := d.AddGoEnvironment()
	filename := env.Codefile(d.CliConf.InputFile())
	fset := token.NewFileSet()

	fmt.Printf("Parsing input file %s\n", filename)
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	LookingForGoGen := 1
	HasGoGenComment := 2

	line := d.Line()
	state := LookingForGoGen
	name := ""
	var docs *FieldAndDoc

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Comment:
			file := fset.File(x.Slash)
			pos := file.Position(x.Slash)
			if pos.Line == line {
				state = HasGoGenComment
			}
		case *ast.TypeSpec:
			name = x.Name.Name
		case *ast.StructType:
			if state == HasGoGenComment {
				docs = NewFieldAndDoc(name)
				d.ProcessFields(docs, x)
				state = LookingForGoGen
			}
		}
		return true
	})
	return docs, nil
}

func (d *DocFinder) ProcessFields(st *FieldAndDoc, x *ast.StructType) {
	st.FieldDoc = make(map[string]string) // gaurantees FieldDoc non-nil
	hasFields := x.Fields != nil && x.Fields.List != nil
	if !hasFields {
		return
	}
	for _, f := range x.Fields.List {
		for _, name := range f.Names {
			comments := make([]string, 0)
			if f.Doc != nil && f.Doc.List != nil {
				for _, doc := range f.Doc.List {
					line := d.stripCommentDelims(doc.Text)
					comments = append(comments, line)
				}
			}
			if f.Comment != nil && f.Comment.List != nil {
				for _, comment := range f.Comment.List {
					line := d.stripCommentDelims(comment.Text)
					comments = append(comments, line)
				}
			}

			// Skip private members
			if strings.ToLower(name.Name)[:1] == name.Name[:1] {
				continue
			}
			st.Add(name.Name, strings.Join(comments, "\n"))
		}
	}
}

func (d *DocFinder) stripCommentDelims(line string) string {
	re := regexp.MustCompile("^\\s*// *")
	return re.ReplaceAllString(line, "")
}

func (d *DocFinder) Run() error {
	structs, err := d.findFieldDocumentation()
	if err != nil {
		return err
	}
	if d.CliConf.HasTemplate() {
		fmt.Println("Rendering struct field and doc to template")
		return d.renderTemplate(structs)
	} else {
		fmt.Println("Rendering struct field and doc to json")
		return d.renderJson(structs)
	}
}

func (d *DocFinder) renderJson(docs *FieldAndDoc) error {
	bb, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		return err
	}

	d.Env.MaybeExit(d.CliConf, nil, string(bb))

	fmt.Printf("Writing output file: %s\n", d.OutputFile())
	file, err := os.Create(d.OutputFile())
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(bb)

	return err
}

func (d *DocFinder) renderTemplate(docs *FieldAndDoc) error {
	fmt.Printf("Reading template file: %s\n", JoinCwd(d.Template()))
	textTemplate, err := ioutil.ReadFile(d.Template())
	if err != nil {
		return err
	}

	fmt.Println("Creating template")
	fmt.Println(string(textTemplate))

	template, err := d.Env.CreateTemplate(string(textTemplate))
	if err != nil {
		fmt.Println(err)
		return err
	}

	d.Env.MaybeExit(d.CliConf, template, "")

	fmt.Printf("Writing output file: %s\n", d.OutputFile())
	file, err := os.Create(d.OutputFile())
	if err != nil {
		return err
	}
	defer file.Close()
	vals := d.Env.ToMap()
	vals["data"] = docs
	template.Execute(file, vals)
	return nil
}
