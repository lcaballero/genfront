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

func Cwd() string {
	return FatalString(os.Getwd())
}

func JoinCwd(s string) string {
	return filepath.Join(Cwd(), s)
}
