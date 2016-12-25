package elastictable

import (
	"github.com/tgulacsi/wrap"
	"github.com/mattn/go-runewidth"
	"strings"
)

func runeWidth(str string) int {
	return runewidth.StringWidth(str)
}

func wrapString(str string, max int) []string {
	return strings.Split(wrap.String(str, uint(max)), "\n")
}
