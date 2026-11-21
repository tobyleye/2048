package main

var textWhite = "#f9f6f2"
var textBlack = "#776e65"

var emptyCellBg = "#313244"
var boardBg = "#11111B"

var cellColors = map[int]*struct {
	fg string
	bg string
}{

	2: {
		fg: textBlack,
		bg: "#eee4da",
	},
	4: {
		fg: textBlack,
		bg: "#eee1c9",
	},

	8: {
		fg: textBlack,
		bg: "#F3B27A",
	},
	16: {
		bg: "#F69664",
		fg: textWhite,
	},
	32: {
		bg: "#F77C5F",
		fg: textWhite,
	},
	64: {
		bg: "#F75F3B",
		fg: textWhite,
	},
	128: {
		bg: "#EDD073",
		fg: textWhite,
	},
	256: {
		bg: "#EDCC62",
		fg: textWhite,
	},
}
