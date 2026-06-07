package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func generateWave(frame int) []string {
	chars := []rune{'.', ',', '-', '~', '^', '"', '`', '\'', '!', '|', 'I', 'l', '1'}
	lines := make([]string, waveRows)
	for row := 0; row < waveRows; row++ {
		var sb strings.Builder
		for col := 0; col < waveCols; col++ {
			px := float64(col) / float64(waveCols)
			py := float64(row) / float64(waveRows)
			v := math.Sin(px*12+float64(frame)*0.08) * 0.30
			v += math.Sin(px*7-float64(frame)*0.05+1.2) * 0.20
			v += math.Sin(py*8+float64(frame)*0.06) * 0.15
			v += math.Sin((px+py)*9+float64(frame)*0.04) * 0.10
			norm := (v + 0.75) / 1.5
			if norm < 0 {
				norm = 0
			}
			if norm > 1 {
				norm = 1
			}
			sb.WriteRune(chars[int(norm*float64(len(chars)-1))])
		}
		lines[row] = sb.String()
	}
	return lines
}

func fmtDur(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	return fmt.Sprintf("%02d:%02d", int(d.Minutes()), int(d.Seconds())%60)
}

func trunc(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n-3]) + "..."
}

func spinner(f int) string {
	return []string{"|", "/", "-", "\\"}[f%4]
}
