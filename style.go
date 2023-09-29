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

import "github.com/charmbracelet/lipgloss"

// Defines the PurpleClay palette of color shades
var (
	S950 = lipgloss.Color("#07000f")
	S900 = lipgloss.Color("#130027")
	S800 = lipgloss.Color("#1e003f")
	S700 = lipgloss.Color("#280057")
	S600 = lipgloss.Color("#33006f")
	S500 = lipgloss.Color("#3d0087")
	S400 = lipgloss.Color("#48009f")
	S300 = lipgloss.Color("#5712ab")
	S200 = lipgloss.Color("#6725b7")
	S100 = lipgloss.Color("#906ccf")
	S50  = lipgloss.Color("#a980db")
)

// Defines the PurpleClay palette for accent color shades
var (
	Green900 = lipgloss.Color("#166534")
	Green700 = lipgloss.Color("#15803d")
)

var (
	// A defines a PurpleClay themed hyperlink that supports both light and dark terminals
	A = lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Foreground(lipgloss.AdaptiveColor{
			Light: string(S400),
			Dark:  string(S100),
		})

	h = lipgloss.NewStyle().Padding(0, 1).Bold(true).Foreground(lipgloss.Color("#ffffff"))

	// H1 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H1 = h.Copy().
		Background(lipgloss.AdaptiveColor{
			Light: string(S50),
			Dark:  string(S200),
		})

	// H2 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H2 = h.Copy().
		Background(lipgloss.AdaptiveColor{
			Light: string(S100),
			Dark:  string(S300),
		})

	// H3 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H3 = h.Copy().
		Background(lipgloss.AdaptiveColor{
			Light: string(S200),
			Dark:  string(S400),
		})

	// H4 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H4 = h.Copy().
		Background(lipgloss.AdaptiveColor{
			Light: string(S300),
			Dark:  string(S500),
		})

	// H5 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H5 = h.Copy().
		Background(lipgloss.AdaptiveColor{
			Light: string(S400),
			Dark:  string(S600),
		})

	// H6 defines a PurpleClay themed header that is ranked on importance from
	// H1 (most) to H6 (least). Supports both light and dark terminals
	H6 = h.Copy().
		Background(lipgloss.AdaptiveColor{
			Light: string(S500),
			Dark:  string(S700),
		})

	// Mark defines a PurpleClay themed text decoration for highlighting text. Supports
	// both light and dark terminals
	Mark = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.AdaptiveColor{
			Light: string(S50),
			Dark:  string(S800),
		})

	// I defines a PurpleClay themed hyperlink that supports both light and dark terminals
	I = lipgloss.NewStyle().Italic(true)

	// U defines a PurpleClay themed underline text decoration
	U = lipgloss.NewStyle().Underline(true)

	// B defines a PurpleClay themed bold text decoration
	B = lipgloss.NewStyle().Bold(true)

	// S defines a PurpleClay themed strikethrough text decoration
	S = lipgloss.NewStyle().Strikethrough(true)
)

var (
	// Tick defines a PurpleClay themed glyph ✓ that supports both light and dark terminals
	Tick = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{
			Light: string(Green900),
			Dark:  string(Green700),
		}).
		Render("✓ ")
)
