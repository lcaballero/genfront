package process
   
import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestName(t *testing.T) {

	Convey("derive outfile should make proper name", t, func() {
		m := map[string]string{
			"my_struct.go":"my_struct_tomap.go",
			"mine.go":"mine_tomap.go",
			"boom":"boom_tomap.go", // we don't want to write over an existing file
		}
		for k,v := range m {
			actual := deriveOutfile(k)
			So(actual, ShouldEqual, v)
		}
	})
}


