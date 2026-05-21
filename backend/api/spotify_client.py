from requests import get, put, post
from config import BASE_URL
from auth.token_manager import get_access_token

class SpotifyClient:
    def __init__(self):
        pass
    
    def headers(self):
        return {
            "Authorization": f"Bearer {get_access_token()}"
        }

    # PRIVATE_PARTS
    def _get(self, endpoint):
        url  =BASE_URL + endpoint
        response = get(url, headers=self.headers())
        return self._handle_response(response)

    def _post(self, endpoint, data=None):
        url = BASE_URL + endpoint
        response = post(url, headers=self.headers(), json=data)
        return self._handle_response(response)

    def _put(self, endpoint, data=None):
        url = BASE_URL + endpoint
        response = put(url, headers=self.headers(), json=data)
        return self._handle_response(response)
    
    def _handle_response(self, response):
        if response.status_code==204:
            return {"success": True}
        
        if response.status_code>=400:
            return {
                "success": False,
                "status": response.status_code,
                "error": response.text
            }

        try:
            if response.json():
                return response.json()
            return {"success": True, "code": response.status_codw}
        except Exception:
            return {"success": False}