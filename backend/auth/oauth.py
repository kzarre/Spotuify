from flask import Flask, redirect, request
import string
import random
import urllib
import base64
from requests import post
import threading
import time
from auth import token_manager as th
from config import *

auth_event = threading.Event()
access_token = None
refresh_token = None


def generate_auth_url(state):
    scope = "user-read-private user-read-email user-top-read user-read-playback-state user-modify-playback-state user-read-currently-playing streaming playlist-read-private playlist-read-collaborative playlist-modify-private playlist-modify-public user-read-playback-position"
    params = {
                "response_type": "code",
                "client_id": CLIENT_ID,
                "scope": scope,
                "redirect_uri": REDIRECT_URI,
                "state": state,
            }
    auth_url = "https://accounts.spotify.com/authorize?" + urllib.parse.urlencode(params)

    return auth_url
    
def exchange_code_for_token(code):
    auth_string = CLIENT_ID + ":" + CLIENT_SECRET
    auth_bytes = auth_string.encode("utf-8")
    auth_base64 = str(base64.b64encode(auth_bytes), "utf-8")

    url = "https://accounts.spotify.com/api/token"
    data = {
            "code" : code,
            "redirect_uri" : REDIRECT_URI,
            "grant_type" : "authorization_code"
            }
    headers = {
                "Authorization" : "Basic " + auth_base64,
                "Content-Type" : "application/x-www-form-urlencoded"
            }


    result = post(url, headers=headers, data=data)
    json_result = result.json()

    return json_result

def random_str(length):
    return "".join(random.choices(string.ascii_letters + string.digits, k=length))



class Login:
    def __init__(self):
        self.app = Flask(__name__)
        self.setup_routes()
        self.event = threading.Event()


    def setup_routes(self):

        @self.app.route("/login")
        def login():
            state = self.random_str(16)
            auth_url = generate_auth_url(state)
            return redirect(auth_url)


        @self.app.route("/callback")
        def callback():
            code  = request.args.get("code")
            
            json_result = exchange_code_for_token(code)
            
            self.access_token = json_result["access_token"]
            self.refresh_token = json_result["refresh_token"]
            
            self.event.set()
            
            return json_result


    def run(self):
        self.app.run(port=8888)


def Oauth():
    auth_event.wait()

    th.set_tokens(access_token, str(time.time()+3540), refresh_token)