/*
Copyright (c) 2023 Purple Clay

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package theme

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// TableBorder defines a series of characters that are used when rendering the
// border of a table
type TableBorder struct {
	Bottom      string
	BottomJoin  string
	BottomLeft  string
	BottomRight string
	Middle      string
	MiddleJoin  string
	MiddleLeft  string
	MiddleRight string
	MiddleTop   string
	Top         string
	TopJoin     string
	TopLeft     string
	TopRight    string
	Vertical    string
}

var (
	bdr = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{
			Light: string(S800),
			Dark:  string(S600),
		})

	cell = lipgloss.NewStyle().Padding(0, 1)

	// ThinBorder defines a series of characters for rendering a table with a thin border
	//
	// 	┌───────┬───────┬───────┐
	// 	│   1   │   2   │   3   │
	// 	├───────┼───────┼───────┤
	//	│   4   │   5   │   6   │
	// 	└───────┴───────┴───────┘
	ThinBorder = TableBorder{
		Bottom:      "─",
		BottomJoin:  "┴",
		BottomLeft:  "└",
		BottomRight: "┘",
		Middle:      "─",
		MiddleJoin:  "┼",
		MiddleLeft:  "├",
		MiddleRight: "┤",
		MiddleTop:   "┬",
		Top:         "─",
		TopJoin:     "┬",
		TopLeft:     "┌",
		TopRight:    "┐",
		Vertical:    "│",
	}

	// NoBorder should be used if rendering a table without a border
	NoBorder = TableBorder{}
)

// Table supports the rendering of tabular data within a terminal
type Table struct {
	border     TableBorder
	rowHeights []int
	colWidths  []int
	data       [][]string
	top        string
	middle     string
	bottom     string
}

// NewTable creates a table that will dynamically size around its provided
// data.
//
//	data := [][]string{
//		{"City", "Avg. Rainfall", "Avg. Temp"},
//		{"Barclona", "640 mm", "21.2 °C"},
//		{"London", "585 mm", "11 °C"},
//	}
//	theme.NewTable(data, theme.ThinBorder)
//
// Fixed widths can be optionally provided. If only one argument is provided
// all columns will be fixed to the same width. If more than one argument is
// provided, each corresponding columns width will be fixed in turn.
//
//	theme.NewTable(data, theme.ThinBorder, 10)
//	theme.NewTable(data, theme.ThinBorder, 20, 15, 15)
func NewTable(data [][]string, border TableBorder, maxW ...int) *Table {
	t := &Table{
		border:     border,
		rowHeights: []int{},
		colWidths:  []int{},
		data:       data,
	}
	t.maxDimensionsFrom(maxW...)
	t.resetDividers()

	return t
}

func (t *Table) maxDimensionsFrom(maxW ...int) {
	if len(t.data) == 0 || len(t.data[0]) == 0 {
		return
	}

	t.rowHeights = make([]int, len(t.data))
	t.colWidths = make([]int, len(t.data[0]))

	for i, row := range t.data {
		for j, c := range row {
			w, h := lipgloss.Size(cell.Render(c))
			t.colWidths[j] = max(t.colWidths[j], w)
			t.rowHeights[i] = max(t.rowHeights[i], h)
		}
	}

	if len(maxW) == 0 {
		return
	}

	if len(maxW) == 1 {
		for i := 0; i < len(t.colWidths); i++ {
			t.colWidths[i] = maxW[0]
		}
	}

	cols := min(len(t.colWidths), len(maxW))
	for i := 0; i < cols; i++ {
		t.colWidths[i] = max(t.colWidths[i], maxW[i])
	}
}

func (t *Table) resetDividers() {
	t.top = divider(
		t.border.TopLeft,
		t.border.Top,
		t.border.TopJoin,
		t.border.TopRight,
		t.colWidths...,
	)

	t.middle = divider(
		t.border.MiddleLeft,
		t.border.Middle,
		t.border.MiddleJoin,
		t.border.MiddleRight,
		t.colWidths...,
	)

	t.bottom = divider(
		t.border.BottomLeft,
		t.border.Bottom,
		t.border.BottomJoin,
		t.border.BottomRight,
		t.colWidths...,
	)
}

func divider(left, sep, join, right string, cellW ...int) string {
	var d strings.Builder
	for _, mw := range cellW {
		d.WriteString(strings.Repeat(sep, mw) + join)
	}
	mid := strings.TrimRight(d.String(), join)

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		bdr.Render(left),
		bdr.Render(mid),
		bdr.Render(right),
	)
}

func verticalDivider(str string, h int) string {
	d := strings.Repeat(str+"\n", h)
	return bdr.Render(strings.TrimRight(d, "\n"))
}

// String renders the table as a formatted string
func (t *Table) String() string {
	if len(t.data) == 0 {
		return ""
	}

	tblRows := make([]string, 0, len(t.data))

	for i, row := range t.data {
		var tblRow []string

		rowH := t.rowHeights[i]
		vertJoin := verticalDivider(t.border.Vertical, rowH)

		for j, col := range row {
			tblRow = append(tblRow, vertJoin)
			tblRow = append(tblRow, cell.Copy().Width(t.colWidths[j]).Render(col))
		}
		tblRow = append(tblRow, vertJoin)
		tblRows = append(tblRows, lipgloss.JoinHorizontal(lipgloss.Left, tblRow...))
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		t.top,
		strings.Join(tblRows, "\n"),
		t.bottom,
	)
}
