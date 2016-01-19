package datafiles


import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"strconv"
)

func TestDataFile(t *testing.T) {

	Convey("Should find 3x3 grid of numbers", t, func() {
		d, _ := NewCsvData("data", "./comma-sep.csv").Parse()

		So(d.Data[0][0], ShouldEqual, "field1")
		So(d.Data[0][1], ShouldEqual, "field2")
		So(d.Data[0][2], ShouldEqual, "field3")

		var n int64 = 1
		for i := 1; i <= 3; i++ {
			for j := 0; j < 3; j++ {
				So(d.Data[i][j], ShouldEqual, strconv.FormatInt(n, 10))
				n++
			}
		}
	})

	Convey("Should read the file and produce lines", t, func() {
		d, err := NewCsvData("data", "./comma-sep.csv").Parse()
		So(err, ShouldBeNil)
		So(d.Key, ShouldEqual, "data")
		So(d.File, ShouldEqual, "./comma-sep.csv")
	})
}

