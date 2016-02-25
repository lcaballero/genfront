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

	Convey("Test camel case simple", t, func() {
		testInput := []string{"First", "second", "third", "word"}
		expectedOutput := "firstSecondThirdWord"

		testOutput := ToCamelCase(testInput...)

		So(testOutput, ShouldEqual, expectedOutput)
	})

	Convey("Test camel case complex", t, func() {
		testInput := []string{"FIRST", "SeCoNd", "tHIRd", "word"}
		expectedOutput := "firstSecondThirdWord"

		testOutput := ToCamelCase(testInput...)

		So(testOutput, ShouldEqual, expectedOutput)
	})
}
