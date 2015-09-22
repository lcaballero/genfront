package main
   
import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
	"bytes"
	"bufio"
	"strings"
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
{{#each names}}
  <li>By {{this}}</li>
{{/each}}
</ul>
`
		reader := bufio.NewReader(bytes.NewBufferString(s))
		portions := &Portions{}
		err := portions.Read(reader)
		So(err, ShouldBeNil)

		err = portions.Render()
		So(err, ShouldBeNil)
		So(strings.Contains(portions.Rendered, "Batman"), ShouldBeTrue)
		So(strings.Contains(portions.Rendered, "Superman"), ShouldBeTrue)
		So(strings.Contains(portions.Rendered, "Green Lantern"), ShouldBeTrue)
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

	Convey("Initial test", t, func() {
		fmt.Println("Initial Test")
		So(true, ShouldEqual, true)
	})
}


