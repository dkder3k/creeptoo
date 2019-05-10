from src.gronsfeld import make_shift


def test_hello_lowercase_right_2019():
    assert make_shift("hello", 2019, "right") == "jemuq"


def test_hello_lowercase_left_2019():
    assert make_shift("hello", 2019, "left") == "fekcm"


def test_jemuq_lowercase_left_2019():
    assert make_shift("jemuq", 2019, "left") == "hello"


def test_fekcm_lowercase_right_2019():
    assert make_shift("fekcm", 2019, "right") == "hello"


def test_hello_mixedcase_right_2019():
    assert make_shift("HeLlo", 2019, "right") == "JeMuq"


def test_hello_mixedcase_left_2019():
    assert make_shift("HeLlo", 2019, "left") == "FeKcm"


def test_jemuq_mixedcase_left_2019():
    assert make_shift("JeMuq", 2019, "left") == "HeLlo"


def test_fekcm_mixedcase_right_2019():
    assert make_shift("FeKcm", 2019, "right") == "HeLlo"


def test_hello_uppercase_right_2019():
    assert make_shift("HELLO", 2019, "right") == "JEMUQ"


def test_hello_uppercase_left_2019():
    assert make_shift("HELLO", 2019, "left") == "FEKCM"


def test_jemuq_uppercase_left_2019():
    assert make_shift("JEMUQ", 2019, "left") == "HELLO"


def test_fekcm_uppercase_right_2019():
    assert make_shift("FEKCM", 2019, "right") == "HELLO"


def test_wonderful_overflow_9():
    assert make_shift("wonderful", 9, "right") == "fxwmnaodu"
    assert make_shift("wonderful", 9, "left") == "nfeuviwlc"
