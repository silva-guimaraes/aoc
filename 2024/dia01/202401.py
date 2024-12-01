
from collections import Counter

input = list(map(lambda x: int(x), open('01.txt', 'r').read().strip().split()))
a = sorted(input[::2])
b = sorted(input[1::2])

def part1():
    print(sum(abs(a[i] - b[i]) for i in range(len(a))))


def part2():
    c = Counter(b)
    print(sum(c.get(i, 0) * i for i in a))

part1()
part2()
