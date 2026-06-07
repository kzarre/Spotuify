package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Song struct {
	Album    string   `json:"album"`
	Artist   []string `json:"artists"`
	Duration int      `json:"duration_ms"`
	Name     string   `json:"name"`
	Spotify  string   `json:"spotify_url"`
	Success  bool     `json:"success"`
}

func Get_song() (Song, error) {
	url := "http://127.0.0.1:8888/current-song"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return Song{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	var current_song Song
	err = json.NewDecoder(resp.Body).Decode(&current_song)

	if err != nil {
		fmt.Println("error decoding json:", err)
		return Song{}, fmt.Errorf("error decoding json: %w", err)
	}

	return current_song, nil
}

func Next_song() (bool, error) {
	url := "http://127.0.0.1:8888/next"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	if err != nil {
		fmt.Println("error decoding json:", err)
		return false, fmt.Errorf("error decoding json: %w", err)
	}

	return true, nil
}

func Prev_song() (bool, error) {
	url := "http://127.0.0.1:8888/prev"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	if err != nil {
		fmt.Println("error decoding json:", err)
		return false, fmt.Errorf("error decoding json: %w", err)
	}

	return true, nil
}

func Play_song() (bool, error) {
	url := "http://127.0.0.1:8888/play"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	if err != nil {
		fmt.Println("error decoding json:", err)
		return false, fmt.Errorf("error decoding json: %w", err)
	}

	return true, nil
}

func Pause_song() (bool, error) {
	url := "http://127.0.0.1:8888/pause"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err)
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	if err != nil {
		fmt.Println("error decoding json:", err)
		return false, fmt.Errorf("error decoding json: %w", err)
	}

	return true, nil
}
