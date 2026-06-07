package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const eqBars = 60
const waveRows = 13
const waveCols = 36

// ─────────────────────────────────────────────────────────────────────────────
// Messages
// ─────────────────────────────────────────────────────────────────────────────

type tickMsg time.Time
type songFetchedMsg Song
type bgSongFetchedMsg Song // Silent background fetch message
type playDoneMsg bool      // true = now playing, false = now paused
type nextDoneMsg struct{}
type prevDoneMsg struct{}
type apiErrMsg struct{ err error }
type errMsg struct{ err error }

// ─────────────────────────────────────────────────────────────────────────────
// Commands
// ─────────────────────────────────────────────────────────────────────────────

// tickCmd updates the UI animations (EQ bars and wave) every 100ms
func tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// bgTickCmd silently polls the API every 5 seconds to check for song updates
func bgTickCmd() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		song, err := Get_song()
		if err != nil {
			return nil // Silently ignore background API failures to prevent UI stuttering
		}
		return bgSongFetchedMsg(song)
	})
}

func fetchSongCmd() tea.Cmd {
	return func() tea.Msg {
		song, err := Get_song()
		if err != nil {
			return errMsg{err}
		}
		return songFetchedMsg(song)
	}
}

func playCmd(currentlyPlaying bool) tea.Cmd {
	return func() tea.Msg {
		if currentlyPlaying {
			_, err := Pause_song()
			if err != nil {
				return apiErrMsg{err}
			}
			return playDoneMsg(false)
		}
		_, err := Play_song()
		if err != nil {
			return apiErrMsg{err}
		}
		return playDoneMsg(true)
	}
}

func nextCmd() tea.Cmd {
	return func() tea.Msg {
		_, err := Next_song()
		if err != nil {
			return apiErrMsg{err}
		}
		return nextDoneMsg{}
	}
}

func prevCmd() tea.Cmd {
	return func() tea.Msg {
		_, err := Prev_song()
		if err != nil {
			return apiErrMsg{err}
		}
		return prevDoneMsg{}
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Model Definition
// ─────────────────────────────────────────────────────────────────────────────

type model struct {
	song       Song
	loaded     bool
	fetchErr   string
	apiErr     string // transient API error (clears on next action)
	elapsed    time.Duration
	playing    bool
	frame      int
	barHeights []float64
	waveLines  []string
	width      int
	height     int
	prog       progress.Model
}

func newModel() model {
	p := progress.New(
		progress.WithSolidFill("#ff3333"),
		progress.WithoutPercentage(),
	)
	p.Empty = '-'
	p.Full = '='
	p.Width = 60

	bars := make([]float64, eqBars)
	for i := range bars {
		bars[i] = rand.Float64()
	}

	return model{
		prog:       p,
		barHeights: bars,
		waveLines:  generateWave(0),
		width:      180,
		height:     40,
	}
}

func (m model) Init() tea.Cmd {
	// Start both the aggressive UI rendering tick and the relaxed background sync tick
	return tea.Batch(fetchSongCmd(), tickCmd(), bgTickCmd())
}

// ─────────────────────────────────────────────────────────────────────────────
// Update
// ─────────────────────────────────────────────────────────────────────────────

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// ── API responses ──────────────────────────────────────────────────────

	case songFetchedMsg:
		m.song = Song(msg)
		m.loaded = true
		m.elapsed = 0
		m.apiErr = ""

	case bgSongFetchedMsg:
		// Quietly update state only if the song actually changed in the background.
		// Bypasses 'm.loaded = false' so the loading screen never flashes.
		if m.song.Spotify != msg.Spotify || m.song.Name != msg.Name {
			m.song = Song(msg)
			m.elapsed = 0
		}
		// Queue up the next background poll cycle
		return m, bgTickCmd()

	case errMsg:
		m.fetchErr = msg.err.Error()

	case apiErrMsg:
		m.apiErr = msg.err.Error()

	case playDoneMsg:
		m.playing = bool(msg)
		m.apiErr = ""

	case nextDoneMsg:
		m.loaded = false
		m.elapsed = 0
		m.playing = true
		return m, fetchSongCmd()

	case prevDoneMsg:
		m.loaded = false
		m.elapsed = 0
		m.playing = true
		return m, fetchSongCmd()

	// ── Window resize ──────────────────────────────────────────────────────

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// ── Keyboard ───────────────────────────────────────────────────────────

	case tea.KeyMsg:
		switch msg.String() {
		case " ", "p":
			if m.loaded {
				return m, playCmd(m.playing)
			}
		case "n":
			return m, nextCmd()
		case "b":
			return m, prevCmd()
		case "R":
			m.loaded = false
			m.fetchErr = ""
			m.apiErr = ""
			return m, fetchSongCmd()
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	// ── Tick ───────────────────────────────────────────────────────────────

	case tickMsg:
		m.frame++
		if m.playing && m.loaded {
			m.elapsed += 100 * time.Millisecond
			dur := time.Duration(m.song.Duration) * time.Millisecond
			if m.elapsed >= dur {
				m.elapsed = dur
				m.playing = false
			}
		}
		for i := range m.barHeights {
			if m.playing {
				target := math.Abs(math.Sin(float64(m.frame)*0.15 + float64(i)*0.42))
				target += math.Abs(math.Sin(float64(m.frame)*0.07+float64(i)*0.19)) * 0.4
				target = math.Min(target/1.4, 1.0)
				m.barHeights[i] += (target - m.barHeights[i]) * 0.25
			} else {
				m.barHeights[i] *= 0.88
			}
		}
		m.waveLines = generateWave(m.frame)
		return m, tickCmd()
	}

	return m, nil
}
