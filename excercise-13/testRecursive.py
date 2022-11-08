from datetime import datetime


def fibonacci(n):
    if n > 1:
        return fibonacci(n - 2) + fibonacci(n - 1)

    return n


if __name__ == "__main__":
    init = datetime.now()

    fibonacci(42)

    end = datetime.now()
    delta = end - init

    print(int(delta.total_seconds() * 1000))
