from api.player import PlayerAPI
from api.tracks import TrackAPI


class PlaybackService:
    def __init__(self):
        self.player_api = PlayerAPI()
        self.track_api = TrackAPI()

    def get_current_song_details(self):
        playback = self.player_api.get_playback_state()
        if not playback.get("success", True):
            return playback

        if not playback.get("item"):
            return {
                "success": False,
                "error": "No track currently playing"
            }
        
        track_id = playback["item"]["id"]

        track = self.track_api.get_track(track_id)

        return {
            "success": True,
            "name": track["name"],
            "album": track["album"]["name"],
            "artists": [
                artist["name"] for artist in track["artists"]
            ],
            "duration_ms": track["duration_ms"],
            "spotify_url": track["external_urls"]["spotify"]
        }
    
    def pause_song(self):
        return self.player_api.pause_playback()
    
    def play_song(self):
        return self.player_api.play_playback()
    
    def next_song(self):
        return self.player_api.next_song()
    
    def prev_song(self):
        return self.player_api.prev_song()