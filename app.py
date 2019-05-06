from flask import Flask, request
from shift import make_shift


app = Flask(__name__)


@app.route("/caesar", methods=["POST", "GET"])
def caesar():
    if request.method == "POST":
        return make_shift(request.form["text"], int(request.form["key"]), request.form["direction"])
    else:
        return "Hello, Caesar"
