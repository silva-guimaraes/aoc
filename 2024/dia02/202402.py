
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

# todas as sequências seguras estão na segunda metadade do input,
# logo, basta parar na primeira encontrada.
def part1():
    for i, a in enumerate(input):
        if safe(a):
            print(len(input) - i)
            return

# pesquisa binária aqui também funciona
def part1binary():
    low, high = 0, len(input)
    while low < high:
        mid = low + int((high - low) / 2)
        if mid == low:
            break
        if safe(input[mid]):
            high = mid
        else:
            low = mid
    print(len(input) - high)


def part2():
    sum = 0
    for a in input:
        if safe(a):
            sum += 1
            continue
        for i in range(len(a)):
            b = a[:]
            del b[i]
            if safe(b):
                sum += 1
                break
    print(sum)


part1()
part1binary()
part2()
