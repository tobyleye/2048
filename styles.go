package main

var textWhite = "#f9f6f2"
var textBlack = "#776e65"

var emptyCellBg = "#313244"
var emptyCellBgDimmed = "#262637"

var boardBg = "#11111B"

var dimmedTextDark = "#423E44"
var dimmedTextDarker = "#76747D"

type tileColor struct {
	fg         string
	bg         string
	dimmedBg   string
	dimmedText string
}

var tileColors = map[int]tileColor{

	2: {
		fg:         textBlack,
		bg:         "#eee4da",
		dimmedBg:   "#716D73",
		dimmedText: dimmedTextDark,
	},
	4: {
		fg:         textBlack,
		bg:         "#eee1c9",
		dimmedBg:   "#716C6C",
		dimmedText: dimmedTextDark,
	},

	8: {
		fg:         textBlack,
		bg:         "#F3B27A",
		dimmedBg:   "#73594C",
		dimmedText: dimmedTextDarker,
	},
	16: {
		bg:         "#F69664",
		fg:         textWhite,
		dimmedBg:   "#744E44",
		dimmedText: dimmedTextDarker,
	},
	32: {
		bg:         "#F77C5F",
		fg:         textWhite,
		dimmedBg:   "#744442",
		dimmedText: dimmedTextDarker,
	},
	64: {
		bg:         "#F75F3B",
		fg:         textWhite,
		dimmedBg:   "#743834",
		dimmedText: dimmedTextDarker,
	},
	128: {
		bg:         "#EDD073",
		fg:         textWhite,
		dimmedBg:   "#71654A",
		dimmedText: dimmedTextDarker,
	},
	256: {
		bg:         "#EDCC62",
		fg:         textWhite,
		dimmedBg:   "#71603E",
		dimmedText: dimmedTextDarker,
	},
	512: {
		bg:         "#EDC850",
		fg:         textWhite,
		dimmedBg:   "#716037",
		dimmedText: dimmedTextDarker,
	},
	1024: {
		bg:         "#EDC53F",
		fg:         textWhite,
		dimmedBg:   "#715F31",
		dimmedText: dimmedTextDarker,
	},
	2048: {
		bg:         "#EDC22E",
		fg:         textWhite,
		dimmedBg:   "#71602E",
		dimmedText: dimmedTextDarker,
	},
	4096: {
		bg:         "#3C3A32",
		fg:         textWhite,
		dimmedBg:   "#2A2A2F",
		dimmedText: dimmedTextDarker,
	},
}

func getTileColor(tile int, dimmed bool) (string, string) {
	c, ok := tileColors[tile]
	if !ok && tile >= 4096 {
		c = tileColors[4096]
	} else if !ok {
		c = tileColors[2]
	}

	if dimmed {
		return c.dimmedBg, c.dimmedText
	}

	return c.bg, c.fg
}
