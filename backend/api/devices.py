from spotify_client import SpotifyClient


class devicesAPI(SpotifyClient):
    def get_devices(self):
        return self._get("/me/player/devices") 

    def transfer_playback(self):
        data = {"device_ids":[self.device_id]}
        return self._put("/me/player`", data=data)