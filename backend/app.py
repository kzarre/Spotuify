from flask import Flask
from routes.auth_routes import auth_bp
from routes.player_routes import player_bp

app = Flask(__name__)

app.register_blueprint(auth_bp)
app.register_blueprint(player_bp)

if __name__=="__main__":
    app.run(port=8888)