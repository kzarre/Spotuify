from flask import Blueprint
from services.playback_service import PlaybackService


player_bp = Blueprint("player", __name__)
playbackService = PlaybackService()

@player_bp.route("/current-song")
def current_song():
    return playbackService.get_current_song_details()

@player_bp.route("/pause")
def pause_song():
    return playbackService.pause_song()

@player_bp.route("/play")
def play_song():
    return playbackService.play_song()

@player_bp.route("/next")
def next_song():
    return playbackService.next_song()

@player_bp.route("/prev")
def prev_song():
    return playbackService.prev_song()
