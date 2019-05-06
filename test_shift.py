from shift import make_shift


def test_hello_lowercase_right_4():
    assert make_shift("hello", 4, "right") == "lipps"


def test_hello_lowercase_left_1():
    assert make_shift("hello", 1, "left") == "gdkkn"


def test_hello_uppercase_right_4():
    assert make_shift("HELLO", 4, "right") == "LIPPS"


def test_hello_uppercase_left_1():
    assert make_shift("HELLO", 1, "left") == "GDKKN"


def test_hello_mixed_left_0():
    assert make_shift("Hello", 0, "left") == "Hello"


def test_hello_mixed_right_0():
    assert make_shift("Hello", 0, "right") == "Hello"


def test_hello_mixed_left_26():
    assert make_shift("Hello", 26, "left") == "Hello"


def test_hello_mixed_right_26():
    assert make_shift("Hello", 26, "right") == "Hello"


def test_hello_mixed_right_overflow():
    assert make_shift("Hello", 27, "right") == "Ifmmp"
    assert make_shift("Hello", 53, "right") == "Ifmmp"
    assert make_shift("Hello", 54, "right") == "Jgnnq"


def test_hello_mixed_left_overflow():
    assert make_shift("Hello", 27, "left") == "Gdkkn"
    assert make_shift("Hello", 53, "left") == "Gdkkn"
    assert make_shift("Hello", 54, "left") == "Fcjjm"


def test_negative_key():
    assert make_shift("hello", -1, "right") == "gdkkn"
    assert make_shift("hello", -4, "left") == "lipps"
