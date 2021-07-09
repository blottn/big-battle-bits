from flask import current_app, flash, jsonify, make_response, redirect, request, url_for, abort
from flask import Flask

from nacl.signing import VerifyKey
from nacl.exceptions import BadSignatureError

APP_ID="862056701553147965"
PUBLIC_KEY="67b9931e4433f10787aed97c14a30b7296f43a73063e4d37a9c1029b11fd9584"

verify_key = VerifyKey(bytes.fromhex(PUBLIC_KEY))

app = Flask(__name__)

@app.route("/")
def index():
    return "<p>Hello, World!</p>"

@app.route("/bot",methods = ['POST'])
def ok():
    signature = request.headers["X-Signature-Ed25519"]
    timestamp = request.headers["X-Signature-Timestamp"]
    body = request.data.decode("utf-8")
    try:
        verify_key.verify(f'{timestamp}{body}'.encode(), bytes.fromhex(signature))
    except BadSignatureError:
        abort(401, 'invalid request signature')

    if request.json["type"] == 1:
        return jsonify({
            "type": 1
        })
    else:
        return jsonify({
            "type": 4,
            "data": {
                "tts": False,
                "content": "Congrats on sending your command!",
                "embeds": [],
                "allowed_mentions": { "parse": [] }
            }
        })

if __name__ == '__main__':
      app.run(port=6969)
