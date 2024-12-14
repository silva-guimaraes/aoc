
from functools import cache
test = [125, 17]
input = [4189, 413, 82070, 61, 655813, 7478611, 0, 8]

@cache
def recur(stone: int, blink: int):
    if blink == 0:
        return 1
    blink -= 1
    if stone == 0:
        return recur(1, blink)
    s = str(stone)
    l = len(s)
    if l % 2 == 0:
        l = l//2
        return recur(int(s[l:]), blink) + recur(int(s[:l]), blink)
    else:
        return recur(stone * 2024, blink)

def part1():
    print(sum(map(lambda x: recur(x, 25), input)))

def part2():
    print(sum(map(lambda x: recur(x, 75), input)))

part1()
part2()
