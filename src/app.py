from flask import Flask, request

app = Flask(__name__)


@app.route("/caesar", methods=["POST", "GET"])
def caesar():
    from src.caesar import make_shift
    if request.method == "POST":
        return make_shift(request.form["text"], int(request.form["key"]), request.form["direction"])
    else:
        return "Hello, Caesar"


@app.route("/gronsfeld", methods=["POST", "GET"])
def gronsfeld():
    from src.gronsfeld import make_shift
    if request.method == "POST":
        return make_shift(request.form["text"], int(request.form["key"]), request.form["direction"])
    else:
        return "Hello, Gronsfeld"
