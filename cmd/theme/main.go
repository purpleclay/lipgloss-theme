package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	theme "github.com/purpleclay/lipgloss-theme"
)

const loremIpsum = "Lorem ipsum dolor sit amet"

func main() {
	out := lipgloss.JoinVertical(
		lipgloss.Top,
		palette(),
		"",
		typogrpahy(),
		"",
		glyphs(),
		"",
		tables(),
	)
	fmt.Fprint(os.Stdout, lipgloss.NewStyle().Margin(2, 2).Render(out))
}

func palette() string {
	colorCell := lipgloss.NewStyle().Height(3).Width(12)
	labelCell := lipgloss.NewStyle().Width(12).AlignHorizontal(lipgloss.Center)

	colors := []string{
		colorCell.Background(theme.S50).Render(),
		colorCell.Background(theme.S100).Render(),
		colorCell.Background(theme.S200).Render(),
		colorCell.Background(theme.S300).Render(),
		colorCell.Background(theme.S400).Render(),
		colorCell.Background(theme.S500).Render(),
		colorCell.Background(theme.S600).Render(),
		colorCell.Background(theme.S700).Render(),
		colorCell.Background(theme.S800).Render(),
		colorCell.Background(theme.S900).Render(),
		colorCell.Background(theme.S950).Render(),
	}
	labels := []string{
		labelCell.Render("50"),
		labelCell.Render("100"),
		labelCell.Render("200"),
		labelCell.Render("300"),
		labelCell.Render("400"),
		labelCell.Render("500"),
		labelCell.Render("600"),
		labelCell.Render("700"),
		labelCell.Render("800"),
		labelCell.Render("900"),
		labelCell.Render("950"),
	}

	table.New().Border(lipgloss.Border{}).Row(colors...).String()

	return lipgloss.JoinVertical(
		lipgloss.Top,
		theme.H1.Render("Monochromatic Palette"),
		theme.NewTable([][]string{colors, labels}).
			Dividers(false).
			Collapsed(true).
			String(),
	)
}

func typogrpahy() string {
	data := [][]string{
		{"h1", theme.H1.Render(loremIpsum), "", "u", theme.U.Render(loremIpsum)},
		{"h2", theme.H2.Render(loremIpsum), "", "i", theme.I.Render(loremIpsum)},
		{"h3", theme.H3.Render(loremIpsum), "", "b", theme.B.Render(loremIpsum)},
		{"h4", theme.H4.Render(loremIpsum), "", "s", theme.S.Render(loremIpsum)},
		{"h5", theme.H5.Render(loremIpsum), "", "a", theme.A.Render(loremIpsum)},
		{"h6", theme.H6.Render(loremIpsum), "", "mark", strings.Replace(loremIpsum, "dolor", theme.Mark.Render("dolor"), 1)},
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		theme.H2.Render("Typography"),
		theme.NewTable(data).Collapsed(true).Dividers(false).Widths(6, 28, 6, 6, 28).String(),
	)
}

func glyphs() string {
	data := [][]string{
		{"tick", theme.Tick + " " + loremIpsum},
		{"cross", theme.Cross + " " + loremIpsum},
		{"bang", theme.Bang + " " + loremIpsum},
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		theme.H3.Render("Glyphs"),
		theme.NewTable(data).Collapsed(true).Dividers(false).Widths(6, 28).String(),
	)
}

func tables() string {
	tbl := [][]string{{"1", "2", "3"}, {"4", "5", "6"}}

	thinBorder := theme.NewTable(tbl).
		Border(theme.ThinBorder).
		Widths(10).
		HorizontalAlignments(lipgloss.Left, lipgloss.Center, lipgloss.Right).
		String()

	thickBorder := theme.NewTable(tbl).
		Border(theme.ThickBorder).
		Widths(10).
		HorizontalAlignments(lipgloss.Left, lipgloss.Center, lipgloss.Right).
		String()

	roundedBorder := theme.NewTable(tbl).
		Border(theme.RoundedThinBorder).
		Widths(10).
		HorizontalAlignments(lipgloss.Left, lipgloss.Center, lipgloss.Right).
		String()

	doubleBorder := theme.NewTable(tbl).
		Border(theme.DoubleBorder).
		Widths(10).
		HorizontalAlignments(lipgloss.Left, lipgloss.Center, lipgloss.Right).
		String()

	labelCell := lipgloss.NewStyle().AlignVertical(lipgloss.Center)
	data := [][]string{
		{
			labelCell.Height(lipgloss.Height(thinBorder)).Render("thin"),
			thinBorder,
			labelCell.Height(lipgloss.Height(thickBorder)).Render("thick"),
			thickBorder,
		},
		{
			labelCell.Height(lipgloss.Height(roundedBorder)).Render("rounded"),
			roundedBorder,
			labelCell.Height(lipgloss.Height(doubleBorder)).Render("double"),
			doubleBorder,
		},
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		theme.H4.Render("Tables"),
		theme.NewTable(data).String(),
	)
}
