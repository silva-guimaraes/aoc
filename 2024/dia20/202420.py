
import heapq
from itertools import combinations

input = open('20.txt').read().strip().split()

start: tuple[int,int] = (-1, -1)
end: tuple[int,int] = (-1, -1)

# gscore = {}
for i in range(len(input)):
    for j in range(len(input[0])):
        # gscore[(i,j)] = math.inf
        if input[i][j] == 'S':
            start = (i,j)
        elif input[i][j] == 'E':
            end = (i,j)

queue: list[tuple[int, int]] = [start]

path = []

while queue:
    i, j = heapq.heappop(queue)

    path.append((i,j))
    if (i, j) == end:
        path = path
        break

    states = [
        (i+1, j, ),
        (i,   j+1),
        (i-1, j, ),
        (i,   j-1),
    ]

    for s in states:
        (ni, nj) = s
        if ni < 0 or ni >= len(input) or nj < 0 or nj >= len(input[0]):
            continue
        if input[ni][nj] == '#':
            continue
        if (ni,nj) in path:
            continue
        queue.append(s)


path.reverse()
# sum = 0
# for c, (i,j) in enumerate(path):
#     state = [
#         (i+2, j  ),
#         (i,   j+2),
#         (i-2, j  ),
#         (i,   j-2),
#     ]
#     for (ni, nj) in state:
#         if ni < 0 or ni >= len(input) or nj < 0 or nj >= len(input[0]):
#             continue
#         try:
#             idx = path.index((ni,nj))
#             if idx-c-1 >= 100:
#                 sum += 1
#         except ValueError:
#             continue
# print(sum)

sum = 0
for ca, (ia, ja) in enumerate(path):
    for cb, (ib, jb) in enumerate(path[ca:]):
        if abs(ia-ib) + abs(ja-jb) < 20 and cb >= 100:
            sum += 1 

print(sum)
