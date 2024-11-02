package theme

var (
	// NoBorder should be used if rendering a table without a border
	NoBorder = TableBorder{}

	// HiddenBorder defines a series of characters for rendering a table with a hidden border,
	// this is different from [NoBorder] as the table will correctly align with all other styles
	HiddenBorder = TableBorder{
		Bottom:      " ",
		BottomJoin:  " ",
		BottomLeft:  " ",
		BottomRight: " ",
		Middle:      " ",
		MiddleJoin:  " ",
		MiddleLeft:  " ",
		MiddleRight: " ",
		MiddleTop:   " ",
		Top:         " ",
		TopJoin:     " ",
		TopLeft:     " ",
		TopRight:    " ",
		Vertical:    " ",
	}

	// ThinBorder defines a series of characters for rendering a table with a thin border
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

	// RoundedThinBorder defines a series of characters for rendering a table with a thin border and rounded edges
	RoundedThinBorder = TableBorder{
		Bottom:      "─",
		BottomJoin:  "┴",
		BottomLeft:  "╰",
		BottomRight: "╯",
		Middle:      "─",
		MiddleJoin:  "┼",
		MiddleLeft:  "├",
		MiddleRight: "┤",
		MiddleTop:   "┬",
		Top:         "─",
		TopJoin:     "┬",
		TopLeft:     "╭",
		TopRight:    "╮",
		Vertical:    "│",
	}

	// ThickBorder defines a series of characters for rendering a table with a thick border
	ThickBorder = TableBorder{
		Bottom:      "━",
		BottomJoin:  "┻",
		BottomLeft:  "┗",
		BottomRight: "┛",
		Middle:      "━",
		MiddleJoin:  "╋",
		MiddleLeft:  "┣",
		MiddleRight: "┫",
		MiddleTop:   "┳",
		Top:         "━",
		TopJoin:     "┳",
		TopLeft:     "┏",
		TopRight:    "┓",
		Vertical:    "┃",
	}

	// DoubleBorder defines a series of characters for rendering a table with a double border
	DoubleBorder = TableBorder{
		Bottom:      "═",
		BottomJoin:  "╩",
		BottomLeft:  "╚",
		BottomRight: "╝",
		Middle:      "═",
		MiddleJoin:  "╬",
		MiddleLeft:  "╠",
		MiddleRight: "╣",
		MiddleTop:   "╦",
		Top:         "═",
		TopJoin:     "╦",
		TopLeft:     "╔",
		TopRight:    "╗",
		Vertical:    "║",
	}
)
