import keyring

tokens = {
    "access_token": None,
    "refresh_token": None,
    "expires_at": None
}

def set_tokens(access_token, expires_at, refresh_token):
   tokens["access_token"] = access_token
   tokens["refresh_token"] = refresh_token
   tokens["expires_at"] = expires_at 

def get_access_token():
    return tokens["access_token"]

def get_refresh_token():
    return tokens["refresh_token"]

def get_expiry():
    return tokens["expires_at"]

def reset():
    keyring.delete_password("spotify_app", "access_token")
    keyring.delete_password("spotify_app", "expiry")
    keyring.delete_password("spotify_app", "refresh_token")

if __name__ == "__main__":
    reset()