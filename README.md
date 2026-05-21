# Spotuify

A terminal-based Spotify controller built with Go and Flask.

Spotuify combines a Flask backend with a Go TUI frontend to create a lightweight Spotify client directly inside the terminal.
The project uses the Spotify Web API for authentication and playback control, while the frontend is being built using Bubble Tea, Bubbles, and Lip Gloss.

> Currently implemented features:
>
> * Spotify OAuth authentication
> * Play / Pause control
> * Shuffle toggle
> * Current song information
> * Modular Flask backend architecture

---

## Tech Stack

### Backend

* Python
* Flask
* Requests
* Spotify Web API
* OAuth2

### Frontend (WIP)

* Go
* Bubble Tea
* Bubbles
* Lip Gloss

---

## Features

* Authenticate with Spotify
* Control playback from terminal
* Toggle shuffle mode
* View currently playing track
* Modular backend structure
* REST API based architecture

---

## Project Structure

```txt
backend/
│   .env
│   .gitignore
│   app.py
│   config.py
│
├───api
│   │   devices.py
│   │   player.py
│   │   spotify_client.py
│   │   tracks.py
│
├───auth
│   │   oauth.py
│   │   token_manager.py
│
├───cache
│
├───routes
│   │   auth_routes.py
│   │   player_routes.py
│
├───services
│   │   playback_service.py
│   │   track_service.py
│
├───utils
```

---

## Architecture

```txt
Go CLI/TUI
    ↓
Flask REST API
    ↓
Spotify Web API
```

The Go frontend communicates with the Flask backend, which handles authentication, token management, and requests to Spotify.

---

## Setup

## 1. Clone the repository

```bash
git clone <your-repo-url>
cd spotuify
```

---

## 2. Create a Spotify Application

Go to the Spotify Developer Dashboard:

[Spotify Developer Dashboard](https://developer.spotify.com/dashboard?utm_source=chatgpt.com)

Create an app and obtain:

* Client ID
* Client Secret

Add a redirect URI such as:

```txt
http://127.0.0.1:5000/callback
```

---

## 3. Configure Environment Variables

Create a `.env` file inside `backend/`

```env
CLIENT_ID=your_client_id
CLIENT_SECRET=your_client_secret
REDIRECT_URI=http://127.0.0.1:5000/callback
```

---

## 4. Install Dependencies

```bash
pip install flask requests python-dotenv
```

---

## 5. Run the Backend

```bash
cd backend
python app.py
```

The Flask server should start locally.

---

## Current API Capabilities

### Authentication

```http
GET /login
GET /callback
```

### Playback Controls

```http
PUT /play
PUT /pause
PUT /shuffle
```

### Track Information

```http
GET /current-track
```

---

## Frontend Status

The frontend is currently under development.

Planned interface:

* Interactive terminal UI
* Real-time playback updates
* Device selection
* Track progress bar
* Keyboard shortcuts
* Search support

---

## Future Plans

* Full Bubble Tea frontend
* Search songs/playlists
* Queue management
* Better caching
* Lyrics support

---

## Notes

Spotify Premium is required for playback control endpoints provided by the Spotify Web API.

---
