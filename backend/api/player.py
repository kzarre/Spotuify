from api.spotify_client import SpotifyClient


class PlayerAPI(SpotifyClient):
    def pause_playback(self):
        return self._put("/me/player/pause")

    def play_playback(self):
        return self._put("/me/player/play")
            
    def next_song(self):
        return self._post("/me/player/next")

    def prev_song(self):
        return self._post("/me/player/previous")

    def set_seek(self, ms):
        return self._post(f"/me/player/seek?position_ms={ms}")
        
    def set_suffle(self, state):
        return self._put(f"/me/player/shuffle?state={state}")

    def set_repeat(self, state):
        return self._put(f"/me/player/repeat?state={state}")
        
    def get_playback_state(self):
        return self._get("/me/player")