
from more_itertools import windowed
from functools import cache
import math
import heapq
# import re
# input = """
# 029A
# 980A
# 179A
# 456A
# 379A
# """
input = """
985A
540A
463A
671A
382A
"""
input = input.split()


directions = ['^', '>', 'v', '<']
numpad: dict[str,list[None | str]] = {
    '9': [None, None, '6', '8'],
    '8': [None, '9', '5', '7'],
    '7': [None, '8', '4', None],
    '6': ['9', None, '3', '5'],
    '5': ['8', '6', '2', '4'],
    '4': ['7', '5', '1', None],
    '3': ['6', None, 'A', '2'],
    '2': ['5', '3', '0', '1'],
    '1': ['4', '2', None, None],
    '0': ['2', 'A', None, None],
    'A': ['3', None, None, '0'],
}

keypad = {
    '^': [None, 'A', 'v', None],
    '>': ['A', None, None, 'v'],
    '<': [None, 'v', None, None],
    'v': ['^', '>', None, '<'],
    'A': [None, None, '>', '^'],
}

# @cache
def goto(pad: dict, start: str, end: str):
    queue: list[tuple[int, str, list[str], set]] = [(0, start, list(), set())]
    shortests: list[str] = []
    while queue:
        c, at, foo, visited = heapq.heappop(queue)
        if at == end:
            t = ''.join(foo[::-1]) + 'A'
            if len(shortests) == 0 or len(shortests[0]) == len(t):
                shortests.append(t)
            continue

        for d, to in enumerate(pad[at]):
            if to is None:
                continue
            if to in visited:
                continue
            heapq.heappush(queue, (c+1, to, [directions[d], *foo], visited | {at}))
    return shortests

# def my_join(ll: list[list[str]], accum: str = "") -> list[str]:
#     if not ll:
#         return [accum]
#     sum = []
#     for l in ll[0]:
#         sum += my_join(ll[1:], accum + l)
#     return sum
# 
# 
# sum = 0
# for i in input:
#     minmin = math.inf
#     s1 = my_join([goto(numpad, a, b) 
#         for a, b in windowed(list('A' + i), 2) if not a is None and not b is None])
#     for s in s1:
#         s2 = my_join([goto(keypad, a, b) 
#             for a, b in windowed(list('A' + s), 2) if not a is None and not b is None])
#         s2 = list(set(s2))
#         for s in s2:
#             p  = [goto(keypad, a, b) 
#                 for a, b in windowed(list('A' + s), 2) if not a is None and not b is None]
#             s3 = my_join(p)
#             minmin = min(minmin, min(map(len, s3)))
# 
#     print(minmin, '*', int(i[:3]))
#     sum += int(i[:3]) * minmin
# print(sum)

def solve(ll: list[list[str]], accum: str = "") -> list[str]:
    if not ll:
        return [accum]
    sum = []
    for l in ll[0]:
        sum += solve(ll[1:], accum + l)
    return sum

def dp(moves: str, i: int):
    if i == 0:
        return min(map(len, moves))

    minmin = math.inf
    for a, b in windowed(list('A' + moves), 2):
        if a is None or b is None: raise
        minmin = min(minmin, goto(numpad, a, b))
    return minmin

print(dp(p))

