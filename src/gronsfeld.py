LEFT, RIGHT = -1, 1


def make_shift(plain_text, key, direction):
    key = str(key)
    direction = LEFT if direction.lower() == "left" else RIGHT
    char_array = list(plain_text)
    for i, c in enumerate(char_array):
        if ('a' <= c <= 'z') or ('A' <= c <= 'Z'):
            shifted_c = chr(ord(c) + int(key[i % len(key)]) * direction)
            if (shifted_c < 'a' <= c) or (shifted_c < 'A'):
                shifted_c = chr(ord(shifted_c) + 26)
            elif shifted_c > 'z' or (shifted_c > 'Z' >= c):
                shifted_c = chr(ord(shifted_c) - 26)
            char_array[i] = shifted_c
    return "".join(char_array)
