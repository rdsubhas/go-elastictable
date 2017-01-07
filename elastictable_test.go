package elastictable

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var dummyHeaders = []string{"h1", "h2", "h3"}

var widthTests = []struct {
	term     int
	given    []int
	expected []int
}{
	{20, []int{5, 10, 15}, []int{3, 3, 3}},
	{40, []int{10, 10, 10}, []int{10, 10, 10}},
	{50, []int{10, 10, 20}, []int{10, 10, 20}},
	{60, []int{10, 25, 70}, []int{10, 14, 26}},
	{60, []int{10, 70, 25}, []int{10, 26, 14}},
	{60, []int{70, 25, 10}, []int{26, 14, 10}},
	{100, []int{50, 50, 50}, []int{30, 30, 30}},
}

func TestOptimizedWidths(t *testing.T) {
	for _, tt := range widthTests {
		os.Setenv(ENV_TERM_WIDTH, fmt.Sprintf("%v", tt.term))
		et := NewElasticTable(dummyHeaders)
		et.AddRow(dummyRow(tt.given...))
		actual := et.optimizedWidths()
		assert.Equal(t, tt.expected, actual)
	}
	os.Setenv(ENV_TERM_WIDTH, "")
}

func TestTableOutput(t *testing.T) {
	given := []string{"column1-10", "column2-10", "column3-10"}
	expected := []string{
		" h1         | h2         | h3         ",
		"------------+------------+------------",
		" column1-10 | column2-10 | column3-10 ",
		"",
	}

	os.Setenv(ENV_TERM_WIDTH, "40")
	buf := new(bytes.Buffer)
	et := NewElasticTable(dummyHeaders)
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

	os.Setenv(ENV_TERM_WIDTH, "40")
	buf := new(bytes.Buffer)
	et := NewElasticTable(dummyHeaders)
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
