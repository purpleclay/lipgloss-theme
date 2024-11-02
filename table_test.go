package theme_test

import (
	"os"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/exp/golden"
	"github.com/muesli/termenv"
	theme "github.com/purpleclay/lipgloss-theme"
)

// Thanks ChatGPT!
// The Madness Rating is an interpretation of each character's level of psychological instability
// and how unpredictable or mentally unhinged they are portrayed
var data = [][]string{
	{"Name", "Sex", "Distinguishing Features", "Madness Rating"},
	{"The Joker", "Male", "Clown-like appearance, green hair, pale skin, psychopathic smile", "10"},
	{"Harley Quinn", "Female", "Clown-like appearance, mallet weapon, acrobatic and unpredictable", "9"},
	{"Two-Face", "Male", "Half-burned face, split personality (Harvey Dent and Two-Face)", "8"},
	{"Scarecrow", "Male", "Wears a scarecrow mask, uses fear toxins to manipulate victims", "8"},
	{"Mad Hatter", "Male", "Obsession with Alice in Wonderland, mind-control technology", "8"},
	{"Riddler", "Male", "Obsession with riddles, green suit with question marks", "7"},
}

func TestMain(m *testing.M) {
	// Strip all color related to the theme as it is breaking golden tests
	lipgloss.SetColorProfile(termenv.Ascii)
	code := m.Run()
	os.Exit(code)
}

func TestTableBorder(t *testing.T) {
	tests := []struct {
		name   string
		border theme.TableBorder
	}{
		{
			name:   "NoBorder",
			border: theme.NoBorder,
		},
		{
			name:   "HiddenBorder",
			border: theme.HiddenBorder,
		},
		{
			name:   "ThinBorder",
			border: theme.ThinBorder,
		},
		{
			name:   "ThickBorder",
			border: theme.ThickBorder,
		},
		{
			name:   "RoundedBorder",
			border: theme.RoundedThinBorder,
		},
		{
			name:   "DoubleBorder",
			border: theme.DoubleBorder,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tbl := theme.NewTable(data).Border(tt.border)

			golden.RequireEqual(t, []byte(tbl.String()))
		})
	}
}

func TestTableNoDividers(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data).
		Border(theme.ThinBorder).
		Dividers(false)

	golden.RequireEqual(t, []byte(tbl.String()))
}

func TestTableCollapsed(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data).
		Collapsed(true)

	golden.RequireEqual(t, []byte(tbl.String()))
}

func TestTableWidths(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data).
		Border(theme.ThinBorder).
		Widths(8, 6, 30, 10)

	golden.RequireEqual(t, []byte(tbl.String()))
}

func TestTableHorizontalAlignments(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data).
		Border(theme.ThinBorder).
		HorizontalAlignments(lipgloss.Left, lipgloss.Center, lipgloss.Left, lipgloss.Right)

	golden.RequireEqual(t, []byte(tbl.String()))
}

func TestTableVerticalAlignments(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data).
		Border(theme.ThinBorder).
		VerticalAlignments(lipgloss.Top, lipgloss.Center, lipgloss.Bottom, lipgloss.Top)

	golden.RequireEqual(t, []byte(tbl.String()))
}
