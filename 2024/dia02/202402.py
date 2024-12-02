
from itertools import pairwise


input = open('02.txt', 'r').read().strip().split('\n')
input = [list(map(int, i.split())) for i in input]

def safe(i):
    a = list(map(lambda x: x[1] - x[0], pairwise(i)))
    diff = map(lambda x: 0 < abs(x) < 4, a)
    asc = map(lambda x: x > 0, a)
    dsc = map(lambda x: x < 0, a)
    if  (all(asc) or all(dsc)) and all(diff):
        return True
    return False

def part1():
    sum = 0
    for i in input:
        if safe(i):
            sum += 1

    print(sum)


def part2():
    sum = 0
    for a in input:
        for i in range(len(a)):
            b = a[:]
            del b[i]
            if safe(b):
                sum += 1
                break
    print(sum)


part1()
part2()
