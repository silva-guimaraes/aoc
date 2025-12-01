
import re
import math
from more_itertools import grouper
import heapq


def simulate(first: list[tuple[int,...]], size: int):
    queue = [(0, 0, 0)]
    gscore = {}
    for i in range(size+1):
        for j in range(size+1):
                gscore[(i,j)] = math.inf

    while queue:
        c, i, j = heapq.heappop(queue)

        if (i, j) == (size, size):
            return c

        states = [
            (c+1, i+1, j),
            (c+1, i, j+1),
            (c+1, i-1, j),
            (c+1, i, j-1),
        ]
        for s in states:
            (tentative, ni, nj) = s
            if ni < 0 or ni > size or nj < 0 or nj > size:
                continue
            if (nj, ni) in first: continue
            if tentative < gscore[(ni, nj)]:
                gscore[(ni,nj)] = tentative
                heapq.heappush(queue, s)
    return None


input = list(grouper(map(int, re.findall(r'\d+', open('18.txt', 'r').read().strip())), 2))
size = 70

def part1():
    print(simulate(input[:1024], size))

def part2():
    low = 0 
    high = len(input)
    while low < high:
        if high - low <= 1:
            print(f'{input[low][0]},{input[low][1]}')
            break
        mid = low + ((high - low) // 2)
        if simulate(input[:mid], size) is None:
            high = mid
        else:
            low = mid

part1()
part2()
