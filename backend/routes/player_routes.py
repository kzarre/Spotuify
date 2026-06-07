from flask import Blueprint, request
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

@player_bp.route("/shuffle")
def set_shuffle():
    state = request.args.get('q') 
    return playbackService.set_shuffle(state)

@player_bp.route("/repeat")
def set_repeat():
    state = request.args.get('q') 
    return playbackService.set_repeat(state)

@player_bp.route("/seek")
def set_seek():
    state = request.args.get('q') 
    return playbackService.set_seek(state)
