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
			Light: string(S400),
			Dark:  string(S200),
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
	dividers   bool
	collapsed  bool
}

// NewTable creates a table that will dynamically size around its provided
// data. By default the table will be rendered without any border
//
//	data := [][]string{
//		{"City", "Avg. Rainfall", "Avg. Temp"},
//		{"Barclona", "640 mm", "21.2 °C"},
//		{"London", "585 mm", "11 °C"},
//	}
//	theme.NewTable(data)
func NewTable(data [][]string) *Table {
	t := &Table{
		border:     NoBorder,
		rowHeights: []int{},
		colWidths:  []int{},
		data:       data,
		dividers:   true,
		collapsed:  false,
	}
	t.maxDimensions()
	t.resetDividers()

	return t
}

func (t *Table) maxDimensions() {
	if len(t.data) == 0 || len(t.data[0]) == 0 {
		return
	}

	cellStyle := cell
	if t.collapsed {
		cellStyle = cellStyle.UnsetPadding()
	}

	t.rowHeights = make([]int, len(t.data))
	t.colWidths = make([]int, len(t.data[0]))

	for i, row := range t.data {
		for j, c := range row {
			w, h := lipgloss.Size(cellStyle.Render(c))
			t.colWidths[j] = max(t.colWidths[j], w)
			t.rowHeights[i] = max(t.rowHeights[i], h)
		}
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

// Border sets the table border
func (t *Table) Border(border TableBorder) *Table {
	t.border = border
	t.resetDividers()
	return t
}

// Widths sets the maximum widths of each colum within the table. If only
// one argument is provided all columns will be fixed to the same width.
// If more than one argument is provided, each corresponding columns width
// will be fixed in turn.
func (t *Table) Widths(w ...int) *Table {
	t.resetMaxWidths(w...)
	t.resetDividers()
	return t
}

func (t *Table) resetMaxWidths(w ...int) {
	if len(w) == 0 {
		return
	}

	if len(w) == 1 {
		for i := 0; i < len(t.colWidths); i++ {
			t.colWidths[i] = w[0]
		}
	}

	cols := min(len(t.colWidths), len(w))
	for i := 0; i < cols; i++ {
		t.colWidths[i] = max(t.colWidths[i], w[i])
	}
}

// Dividers controls whether a row divider should be rendered
// between all table rows
func (t *Table) Dividers(on bool) *Table {
	t.dividers = on
	return t
}

// Collapsed controls whether all internal padding within the
// table should be removed
func (t *Table) Collapsed(on bool) *Table {
	t.collapsed = on
	t.maxDimensions()
	return t
}

// String renders the table as a formatted string
func (t *Table) String() string {
	if len(t.data) == 0 {
		return ""
	}

	cellStyle := cell
	if t.collapsed {
		cellStyle = cellStyle.UnsetPadding()
	}

	var tblRows []string
	for i, row := range t.data {
		var tblRow []string

		rowH := t.rowHeights[i]
		vertJoin := verticalDivider(t.border.Vertical, rowH)

		for j, col := range row {
			tblRow = append(tblRow, vertJoin, cellStyle.Width(t.colWidths[j]).Render(col))
		}
		tblRow = append(tblRow, vertJoin)
		tblRows = append(tblRows, lipgloss.JoinHorizontal(lipgloss.Left, tblRow...))
		if t.dividers && i < len(t.data)-1 {
			tblRows = append(tblRows, t.middle)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		t.top,
		strings.Join(tblRows, "\n"),
		t.bottom,
	)
}
