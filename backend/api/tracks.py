from api.spotify_client import SpotifyClient

class TrackAPI(SpotifyClient):
    def get_track(self, song_id):
        return self._get(f"/tracks/{song_id}")