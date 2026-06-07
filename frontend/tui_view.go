package main

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	mg := lipgloss.NewStyle().Foreground(colMagenta).Background(colBlack).Bold(true)
	grn := lipgloss.NewStyle().Foreground(colGreen).Background(colBlack)
	dim := lipgloss.NewStyle().Foreground(colGray).Background(colBlack)
	red := lipgloss.NewStyle().Foreground(colRed).Background(colBlack)

	totalW := m.width
	if totalW < 80 {
		totalW = 80
	}

	// 1. Title bar (Spans the entire top)
	titleBar := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(colMagenta).
		Background(colBlack).
		Width(totalW - 2).
		Align(lipgloss.Center).
		Render(mg.Render("♪ SpoTUIfy"))

	// Loading / error states
	if m.fetchErr != "" {
		body := panel(
			red.Render("[!] "+m.fetchErr)+"\n\n"+dim.Render("Press R to retry  |  q to quit"),
			totalW-4,
		)
		return lipgloss.JoinVertical(lipgloss.Left, titleBar, body)
	}
	if !m.loaded {
		body := panel(
			grn.Render(spinner(m.frame)+"  Fetching current track..."),
			totalW-4,
		)
		return lipgloss.JoinVertical(lipgloss.Left, titleBar, body)
	}

	// 2. Middle Row Calculations
	// Calculate how wide the right side should be based on the fixed wave width
	waveW := waveCols + 6
	rightColW := totalW - waveW - 2

	wavePanel := renderWavePanel(m)
	infoPanel := renderInfoPanel(m, rightColW)
	ctrlsPanel := renderControls(rightColW)

	// Stack Info and Controls on top of each other for the right column
	rightColumn := lipgloss.JoinVertical(lipgloss.Left, infoPanel, ctrlsPanel)

	// Glue the Left (Wave) and Right (Info/Controls) together horizontally
	middleRow := lipgloss.JoinHorizontal(lipgloss.Top, wavePanel, rightColumn)

	// 3. Bottom Row (EQ Bars span the full width)
	// We pass totalW-2 to the component wrapper inside renderEQPanel below
	eqPanel := renderFullWidthEQPanel(m, totalW-4)

	// 4. Assemble the whole dashboard vertically: Top -> Middle -> Bottom
	return lipgloss.JoinVertical(lipgloss.Left, titleBar, middleRow, eqPanel)
}

func renderFullWidthEQPanel(m model, w int) string {
	const eqHeight = 8 // Dropped slightly to 8 rows so it doesn't eat too much vertical screen space
	green := lipgloss.NewStyle().Foreground(colGreen).Background(colBlack)
	greenDim := lipgloss.NewStyle().Foreground(colGreenDim).Background(colBlack)
	gray := lipgloss.NewStyle().Foreground(colGray).Background(colBlack)

	rows := make([]string, eqHeight)
	for row := 0; row < eqHeight; row++ {
		thresh := 1.0 - float64(row)/float64(eqHeight)
		var sb strings.Builder
		for i, h := range m.barHeights {
			var ch string
			if h >= thresh {
				if thresh > 0.7 {
					ch = green.Render("#")
				} else {
					ch = greenDim.Render("#")
				}
			} else {
				ch = gray.Render(".")
			}
			sb.WriteString(ch)
			if i < len(m.barHeights)-1 {
				sb.WriteString(gray.Render(" "))
			}
		}
		rows[row] = sb.String()
	}
	// Wraps the EQ canvas perfectly into the layout's full available bottom width
	return panel(strings.Join(rows, "\n"), w+2)
}

func renderWavePanel(m model) string {
	green := lipgloss.NewStyle().Foreground(colGreen).Background(colBlack)
	dimG := lipgloss.NewStyle().Foreground(colGreenDim).Background(colBlack)
	var sb strings.Builder
	for i, line := range m.waveLines {
		runes := []rune(line)
		for j, ch := range runes {
			bright := ch == '^' || ch == '|' || ch == '!' || ch == 'I' || ch == '1' || ch == 'l'
			edge := i == 0 || i == waveRows-1 || j == 0 || j == len(runes)-1
			if bright && !edge {
				sb.WriteString(green.Render(string(ch)))
			} else {
				sb.WriteString(dimG.Render(string(ch)))
			}
		}
		if i < waveRows-1 {
			sb.WriteRune('\n')
		}
	}
	return panel(sb.String(), waveCols+4)
}

func renderEQPanel(m model) string {
	const eqHeight = 12
	green := lipgloss.NewStyle().Foreground(colGreen).Background(colBlack)
	greenDim := lipgloss.NewStyle().Foreground(colGreenDim).Background(colBlack)
	gray := lipgloss.NewStyle().Foreground(colGray).Background(colBlack)

	rows := make([]string, eqHeight)
	for row := 0; row < eqHeight; row++ {
		thresh := 1.0 - float64(row)/float64(eqHeight)
		var sb strings.Builder
		for i, h := range m.barHeights {
			var ch string
			if h >= thresh {
				if thresh > 0.7 {
					ch = green.Render("#")
				} else {
					ch = greenDim.Render("#")
				}
			} else {
				ch = gray.Render(".")
			}
			sb.WriteString(ch)
			if i < len(m.barHeights)-1 {
				sb.WriteString(gray.Render(" "))
			}
		}
		rows[row] = sb.String()
	}
	return panel(strings.Join(rows, "\n"), eqBars*2+2)
}

func renderInfoPanel(m model, w int) string {
	mg := lipgloss.NewStyle().Foreground(colMagenta).Background(colBlack)
	grn := lipgloss.NewStyle().Foreground(colGreen).Background(colBlack)
	red := lipgloss.NewStyle().Foreground(colRed).Background(colBlack)
	wh := lipgloss.NewStyle().Foreground(colWhite).Background(colBlack)
	dim := lipgloss.NewStyle().Foreground(colGray).Background(colBlack)

	artists := strings.Join(m.song.Artist, ", ")
	inner := w - 4

	line1 := mg.Render("Now Playing : ") + grn.Render(trunc(m.song.Name, inner-14))
	artistLine := dim.Render("Artist : ") + wh.Render(trunc(artists, inner-10))
	albumLine := dim.Render("Album  : ") + wh.Render(trunc(m.song.Album, inner-10))

	dur := time.Duration(m.song.Duration) * time.Millisecond
	ratio := 0.0
	if dur > 0 {
		ratio = m.elapsed.Seconds() / dur.Seconds()
		if ratio > 1 {
			ratio = 1
		}
	}
	m.prog.Width = inner - 14
	bar := red.Render(fmtDur(m.elapsed)) + " " + m.prog.ViewAs(ratio) + " " + red.Render(fmtDur(dur))

	var statusStr string
	if m.playing {
		statusStr = grn.Render(">> PLAYING")
	} else {
		statusStr = mg.Render("|| PAUSED")
	}

	errLine := ""
	if m.apiErr != "" {
		errLine = red.Render("[!] " + trunc(m.apiErr, inner))
	}

	spotLine := dim.Render(trunc(m.song.Spotify, inner))

	rows := []string{line1, artistLine, albumLine, "", bar, "", statusStr}
	if errLine != "" {
		rows = append(rows, errLine)
	}
	rows = append(rows, "", spotLine)

	return panel(lipgloss.JoinVertical(lipgloss.Left, rows...), w)
}

func renderControls(w int) string {
	mg := lipgloss.NewStyle().Foreground(colMagenta).Background(colBlack).Bold(true)
	grn := lipgloss.NewStyle().Foreground(colGreen).Background(colBlack)
	dim := lipgloss.NewStyle().Foreground(colGray).Background(colBlack)

	keys := []string{
		grn.Render("(p)") + dim.Render(" play/pause"),
		grn.Render("(n)") + dim.Render(" next"),
		grn.Render("(b)") + dim.Render(" prev"),
		grn.Render("(R)") + dim.Render(" refresh"),
		grn.Render("(q)") + dim.Render(" quit"),
	}

	ctrl := mg.Render("Controls") + "\n" + strings.Join(keys, "  ")
	return panel(ctrl, w)
}
