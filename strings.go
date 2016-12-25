package main

import (
	"github.com/tgulacsi/wrap"
	"strings"
)

func runeWidth(str string) int {
	return len(str)
}

func wrapString(str string, max int) []string {
	return strings.Split(wrap.String(str, uint(max)), "\n")
}
