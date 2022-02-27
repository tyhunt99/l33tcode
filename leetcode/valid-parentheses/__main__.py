
def isValid(s: str) -> bool:
    stack = []
    pairs = {
        "{": "}",
        "[": "]",
        "(": ")",
    }
    for char in s:
        # append opening chars
        if char in pairs:
            stack.append(pairs[char])
        # starting with invalid char
        elif len(stack) == 0:
            return False
        # check closing pairs
        elif len(stack) > 0 and stack.pop() != char:
            return False
    return len(stack) == 0

if __name__ == "__main__":
    tests = [
        "()",
        "()[]{}",
        "(]",
        "([}])",
        "[([}])",
        "[",
        "]",
    ]

    for t in tests:
        print(isValid(t))
        print("==============================")
