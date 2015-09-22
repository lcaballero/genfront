package main

import (
	"fmt"
	"strings"
)

func toPascal(s string) string {
	start := strings.ToUpper(s[0:1])
	end := strings.ToLower(s[1:])
	return fmt.Sprintf("%s%s", start, end)
}

