package doctable

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFieldAndDoc(t *testing.T) {

	Convey("commentToString should remove multi-line comment delimiters /* and */", t, func() {
		d := NewFieldAndDoc("conf")
		d.Add("max", "/* and */")
		doc := d.FieldDoc["max"]
		fmt.Println(doc)
		So(doc, ShouldEqual, "and")
	})

	Convey("commentToString should remove line comment start // and trim lead+trailing whitespace", t, func() {
		d := NewFieldAndDoc("conf")
		d.Add("max", "// Here")
		doc := d.FieldDoc["max"]
		fmt.Println(doc)
		So(doc, ShouldEqual, "Here")
	})

	Convey("Adding field should save field with comment", t, func() {
		f := NewFieldAndDoc("name")
		f.Add("field1", "comment1")

		v, ok := f.FieldDoc["field1"]

		So(ok, ShouldBeTrue)
		So(v, ShouldEqual, "comment1")
	})

	Convey("New FieldAndComment should start with given name", t, func() {
		f := NewFieldAndDoc("name")
		So(f.Name, ShouldEqual, "name")
	})
}
