package elastictable

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"strings"
	"bytes"
	"fmt"
	"os"
)

var dummy_headers = []string{"h1", "h2", "h3"}

var widthTests = []struct {
	term int
	given []int
	expected []int
}{
	{ 20, []int{5,10,15}, []int{5,5,5} },
	{ 40, []int{10,10,10}, []int{10,10,11} },
	{ 50, []int{10,10,20}, []int{10,10,21} },
	{ 60, []int{10,25,70}, []int{10,13,28} },
	{ 60, []int{10,70,25}, []int{10,28,13} },
	{ 60, []int{70,25,10}, []int{28,13,10} },
	{ 60, []int{70,10,25}, []int{28,10,13} },
}

func TestOptimizedWidths(t *testing.T) {
	for _, tt := range widthTests {
		os.Setenv(ENV_TERM_WIDTH, fmt.Sprintf("%v", tt.term))
		et := NewElasticTable(dummy_headers)
		et.AddRow(dummyRow(tt.given...))
		actual := et.optimizedWidths()
		assert.Equal(t, tt.expected, actual)
	}
	os.Setenv(ENV_TERM_WIDTH, "")
}

func TestTableOutput(t *testing.T) {
	given := []string{"column1-10", "column2-10", "column3-10"}
	expected := []string{
		" h1         | h2         | h3          ",
		"------------+------------+-------------",
		" column1-10 | column2-10 | column3-10  ",
		"",
	}

	os.Setenv(ENV_TERM_WIDTH, "40")
	buf := new(bytes.Buffer)
	et := NewElasticTable(dummy_headers)
	et.AddRow(given)
	et.Render(buf)
	actual := strings.Split(buf.String(), "\n")

	assert.Equal(t, expected, actual)
}

func TestWrapping(t *testing.T) {
	given := []string{"column1-10", "column2-10column2-20", "column3-10column3-20column3-30"}
	expected := []string{
		" h1         | h2         | h3         ",
		"------------+------------+------------",
		" column1-10 | column2-10 | column3-10 ",
		"            | column2-20 | column3-20 ",
		"            |            | column3-30 ",
		"",
	}

	os.Setenv(ENV_TERM_WIDTH, "30")
	buf := new(bytes.Buffer)
	et := NewElasticTable(dummy_headers)
	et.AddRow(given)
	et.Render(buf)
	actual := strings.Split(buf.String(), "\n")

	assert.Equal(t, expected, actual)
}

func dummyRow(widths ...int) []string {
	row := make([]string, len(widths))
	for i, w := range widths {
		row[i] = strings.Repeat("z", w)
	}
	return row
}