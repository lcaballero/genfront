package maybe
import (
	"log"
	"os"
	"path/filepath"
)


func FatalString(s string, err error) string {
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func JoinCwd(s string) string {
	return filepath.Join(FatalString(os.Getwd()), s)
}
