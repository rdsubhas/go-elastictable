package elastictable

import (
	"github.com/olekukonko/ts"
	"strconv"
	"os"
)

const ENV_TERM_WIDTH = "COLUMNS"
const DEFAULT_TERM_WIDTH = 80

func termWidth() (int) {
	if termWidth, err := strconv.Atoi(os.Getenv(ENV_TERM_WIDTH)); err == nil && termWidth > 0 {
		return termWidth
	} else if termSize, err := ts.GetSize(); err == nil && termSize.Col() > 0 {
		return termSize.Col()
	} else {
		return DEFAULT_TERM_WIDTH
	}
}
