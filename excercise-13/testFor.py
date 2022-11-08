from datetime import datetime


def test_for(n):
    for i in range(n):
        #print(i)
        pass


if __name__ == "__main__":
    init = datetime.now()

    test_for(10000000)

    end = datetime.now()
    delta = end - init

    print(int(delta.total_seconds() * 1000))
