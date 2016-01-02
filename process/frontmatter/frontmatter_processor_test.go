package frontmatter

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFrontmatter(t *testing.T) {

	Convey("Read should partition the file into frontmatter and template", t, func() {
		s := `
---
names:
  - Batman
  - Superman
  - Green Lantern
---
<ul>
{{ range .names }}
  <li>By {{ . }}</li>
{{ end }}
</ul>
`
		reader := bufio.NewReader(bytes.NewBufferString(s))
		portions := &Portions{}
		err := portions.Read(reader)
		So(err, ShouldBeNil)

		tpl, err := portions.CreateTemplate()
		buf := bytes.NewBufferString("")
		tpl.Execute(buf, portions.Settings())
		rendered := buf.String()

		So(err, ShouldBeNil)
		So(strings.Contains(rendered, "Batman"), ShouldBeTrue)
		So(strings.Contains(rendered, "Superman"), ShouldBeTrue)
		So(strings.Contains(rendered, "Green Lantern"), ShouldBeTrue)
	})

	Convey("Read should partition the file into frontmatter and template", t, func() {
		s := `
---
1
---
2
`
		reader := bufio.NewReader(bytes.NewBufferString(s))
		portions := &Portions{}
		err := portions.Read(reader)

		So(err, ShouldBeNil)
		So(portions, ShouldNotBeNil)
		So(portions.FrontMatter, ShouldEqual, "1")
		So(portions.Template, ShouldEqual, "2")
	})
}
