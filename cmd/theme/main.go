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

package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
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
		colorCell.Copy().Background(theme.S50).Render(),
		colorCell.Copy().Background(theme.S100).Render(),
		colorCell.Copy().Background(theme.S200).Render(),
		colorCell.Copy().Background(theme.S300).Render(),
		colorCell.Copy().Background(theme.S400).Render(),
		colorCell.Copy().Background(theme.S500).Render(),
		colorCell.Copy().Background(theme.S600).Render(),
		colorCell.Copy().Background(theme.S700).Render(),
		colorCell.Copy().Background(theme.S800).Render(),
		colorCell.Copy().Background(theme.S900).Render(),
		colorCell.Copy().Background(theme.S950).Render(),
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

	return lipgloss.JoinVertical(
		lipgloss.Top,
		theme.H1.Render("Monochromatic Palette"),
		theme.NewTable([][]string{colors, labels}, theme.NoBorder).String(),
	)
}

func typogrpahy() string {
	data := [][]string{
		{"h1", theme.H1.Render(loremIpsum), "", "u", theme.U.Render(loremIpsum)},
		{"h2", theme.H2.Render(loremIpsum), "", "i", theme.I.Render(loremIpsum)},
		{"h3", theme.H3.Render(loremIpsum), "", "b", theme.B.Render(loremIpsum)},
		{"h4", theme.H4.Render(loremIpsum), "", "s", theme.S.Render(loremIpsum)},
		{"h5", theme.H5.Render(loremIpsum), "", "a", theme.A.Render(loremIpsum)},
		{"h6", theme.H6.Render(loremIpsum), "", "", ""},
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		theme.H2.Render("Typography"),
		theme.NewTable(data, theme.NoBorder, 6, 12, 6, 6, 12).String(),
	)
}

func glyphs() string {
	data := [][]string{
		{"tick", theme.Tick + loremIpsum},
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		theme.H3.Render("Glyphs"),
		theme.NewTable(data, theme.NoBorder, 6, 12).String(),
	)
}

func tables() string {
	tbl := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

	thinBorder := theme.NewTable(tbl, theme.ThinBorder, 6).String()

	labelCell := lipgloss.NewStyle().AlignVertical(lipgloss.Center)
	data := [][]string{
		{labelCell.Copy().Height(lipgloss.Height(thinBorder)).Render("thin"), thinBorder},
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		theme.H4.Render("Tables"),
		theme.NewTable(data, theme.NoBorder, 6, 20).String(),
	)
}
