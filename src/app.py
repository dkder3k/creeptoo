from flask import Flask, request

application = Flask(__name__)


@application.route("/caesar", methods=["POST", "GET"])
def caesar():
    from caesar import make_shift
    if request.method == "POST":
        return make_shift(request.form["text"], int(request.form["key"]), request.form["direction"])
    else:
        return "Hello, Caesar"


@application.route("/gronsfeld", methods=["POST", "GET"])
def gronsfeld():
    from gronsfeld import make_shift
    if request.method == "POST":
        return make_shift(request.form["text"], int(request.form["key"]), request.form["direction"])
    else:
        return "Hello, Gronsfeld"
