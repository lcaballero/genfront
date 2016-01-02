package process

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelpers(t *testing.T) {

	Convey("Initial test", t, func() {
		m := map[string]string{
			"owned_by":   "OwnedBy",
			"created_by": "CreatedBy",
			"":           "",
			"i":          "I",
			"Id":         "Id",
			"id":         "Id",
		}

		for k, exptected := range m {
			actual := SnakeToPascal(k)
			So(actual, ShouldEqual, exptected)
		}
	})
}
