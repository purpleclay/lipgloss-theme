package theme_test

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	theme "github.com/purpleclay/lipgloss-theme"
	"gotest.tools/v3/golden"
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

func TestTableNoBorder(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data)

	golden.Assert(t, lipgloss.NewStyle().Render(tbl.String()), "TestTableNoBorder.golden")
}

func TestTableThinBorder(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data).
		Border(theme.ThinBorder)

	golden.Assert(t, lipgloss.NewStyle().Render(tbl.String()), "TestTableThinBorder.golden")
}

func TestTableNoDividers(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data).
		Border(theme.ThinBorder).
		Dividers(false)

	golden.Assert(t, lipgloss.NewStyle().Render(tbl.String()), "TestTableNoDividers.golden")
}

func TestTableCollapsed(t *testing.T) {
	t.Parallel()
	tbl := theme.NewTable(data).
		Collapsed(true)

	golden.Assert(t, lipgloss.NewStyle().Render(tbl.String()), "TestTableCollapsed.golden")
}

func TestTableColumnAlignment(t *testing.T) {
}

func TestTableMixedColumnAlignment(t *testing.T) {
}
