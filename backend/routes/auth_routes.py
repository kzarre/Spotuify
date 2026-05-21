from flask import Blueprint, redirect, request
from auth.token_manager import get_access_token, set_tokens
from auth.oauth import (
    generate_auth_url,
    exchange_code_for_token,
    random_str
)
import auth.oauth as oauth
import time

auth_bp = Blueprint("auth", __name__)


@auth_bp.route("/login")
def login():
    state = random_str(16)
    auth_url = generate_auth_url(state)

    return redirect(auth_url)

@auth_bp.route("/callback")
def callback():
    code = request.args.get("code")

    json_result = exchange_code_for_token(code)

    oauth.access_token = json_result["access_token"]
    oauth.refresh_token = json_result["refresh_token"]

    set_tokens(json_result["access_token"], time.time()+3540, json_result["refresh_token"])

    return json_result