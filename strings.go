package elastictable

import (
	"github.com/mattn/go-runewidth"
	"github.com/tgulacsi/wrap"
	"strings"
)

func runeWidth(str string) int {
	return runewidth.StringWidth(str)
}

func wrapString(str string, max int) []string {
	return strings.Split(wrap.String(str, uint(max)), "\n")
}
