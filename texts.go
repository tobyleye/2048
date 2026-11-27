package main

import "charm.land/lipgloss/v2"

var cardBg = lipgloss.Color(boardBg)

var keyHintStyle = lipgloss.NewStyle().Background(
	lipgloss.Color("#45475A"),
).Foreground(
	lipgloss.Color("#F9E2AF")).Padding(0, 1)

var gameOver = lipgloss.NewStyle().
	Background(cardBg).
	Padding(1, 0).
	Render(
		lipgloss.JoinVertical(
			lipgloss.Bottom,
			lipgloss.NewStyle().MarginBottom(1).Render(
				lipgloss.PlaceHorizontal(
					gridWidth,
					lipgloss.Center,
					lipgloss.NewStyle().
						Foreground(lipgloss.Color("#F38BA8")).
						Background(cardBg).
						Render("game over"),
					lipgloss.WithWhitespaceStyle(
						lipgloss.NewStyle().
							Background(cardBg),
					),
				),
			),
			// lipgloss.PlaceHorizontal(
			// 	gridWidth,
			// 	lipgloss.Center,
			// 	lipgloss.NewStyle().
			// 		Foreground(lipgloss.Color("#A6E3A1")).
			// 		Background(cardBg).
			// 		Render("no moves left"),
			// 	lipgloss.WithWhitespaceStyle(
			// 		lipgloss.NewStyle().
			// 			Background(cardBg),
			// 	),
			// ),

			lipgloss.PlaceHorizontal(
				gridWidth,
				lipgloss.Center,
				lipgloss.JoinHorizontal(
					lipgloss.Left,
					keyHintStyle.
						Render("r"),
					lipgloss.NewStyle().
						Background(cardBg).
						Render(" new game "),
					keyHintStyle.Render("q"),
					lipgloss.NewStyle().Background(cardBg).Render(" quit"),
				),
				lipgloss.WithWhitespaceStyle(
					lipgloss.NewStyle().
						Background(cardBg),
				),
			),
		),
	)

var congratsMessage = lipgloss.NewStyle().
	Background(cardBg).
	Padding(1, 0).
	Render(
		lipgloss.JoinVertical(
			lipgloss.Bottom,
			lipgloss.NewStyle().MarginBottom(1).Render(
				lipgloss.PlaceHorizontal(
					gridWidth,
					lipgloss.Center,
					lipgloss.NewStyle().
						Foreground(lipgloss.Color("#EDC22E")).
						Background(cardBg).
						Render("Congrats, 2048!"),
					lipgloss.WithWhitespaceStyle(
						lipgloss.NewStyle().
							Background(cardBg),
					),
				),
			),

			lipgloss.PlaceHorizontal(
				gridWidth,
				lipgloss.Center,
				lipgloss.JoinHorizontal(
					lipgloss.Left,
					keyHintStyle.
						Render("c"),
					lipgloss.NewStyle().
						Background(cardBg).
						Render(" continue "),
					keyHintStyle.Render("r"),
					lipgloss.NewStyle().Background(cardBg).Render(" restart "),
				),
				lipgloss.WithWhitespaceStyle(
					lipgloss.NewStyle().
						Background(cardBg),
				),
			),

			// lipgloss.PlaceHorizontal(
			// 	gridWidth,
			// 	lipgloss.Center,
			// 	lipgloss.JoinHorizontal(
			// 		lipgloss.Left,
			// 		keyHintStyle.Render("q"),
			// 		lipgloss.NewStyle().Background(cardBg).Render(" quit"),
			// 	),
			// 	lipgloss.WithWhitespaceStyle(
			// 		lipgloss.NewStyle().
			// 			Background(cardBg),
			// 	),
			// ),
		),
	)
