package elastictable

import (
	"strings"
	"unicode/utf8"

	"github.com/eidolon/wordwrap"
)

func runeWidth(str string) int {
	return utf8.RuneCountInString(str)
}

func wrapString(str string, max int) []string {
	wrapper := wordwrap.Wrapper(max, true)
	return strings.Split(wrapper(str), "\n")
}
