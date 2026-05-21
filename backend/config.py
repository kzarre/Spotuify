import os
from dotenv import load_dotenv

load_dotenv()

CLIENT_ID = os.getenv("CLIENT_ID")
CLIENT_SECRET = os.getenv("CLIENT_SECRET")
REDIRECT_URI = os.getenv("REDIRECT_URI")

TOKEN_URL = "https://accounts.spotify.com/api/token"
BASE_URL = "https://api.spotify.com/v1"