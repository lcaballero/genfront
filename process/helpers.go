package process

import (
	"fmt"
	"strings"
)

func ToPascal(s string) string {
	start := strings.ToUpper(s[0:1])
	end := strings.ToLower(s[1:])
	return fmt.Sprintf("%s%s", start, end)
}

func SnakeToPascal(sk string) string {
	if len(sk) <= 1 {
		return strings.ToUpper(sk)
	}
	parts := strings.Split(sk, "_")
	for i, p := range parts {
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

func ToSymbol(sys string) string {
	parts := strings.Split(sys, " ")
	for i := 0; i < len(parts); i++ {
		parts[i] = ToPascal(parts[i])
	}
	header := strings.Join(parts, "")
	return header
}
