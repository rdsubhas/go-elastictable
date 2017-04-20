package elastictable

import (
	"github.com/tgulacsi/wrap"
	"strings"
	"unicode/utf8"
)

func runeWidth(str string) int {
	return utf8.RuneCountInString(str)
}

func wrapString(str string, max int) []string {
	return strings.Split(wrap.String(str, uint(max)), "\n")
}
