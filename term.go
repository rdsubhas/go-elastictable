package elastictable

import (
	"github.com/olekukonko/ts"
	"os"
	"strconv"
)

const envTermWidth = "COLUMNS"
const defaultTermWidth = 80

func termWidth() int {
	if termWidth, err := strconv.Atoi(os.Getenv(envTermWidth)); err == nil && termWidth > 0 {
		return termWidth
	} else if termSize, err := ts.GetSize(); err == nil && termSize.Col() > 0 {
		return termSize.Col()
	} else {
		return defaultTermWidth
	}
}
