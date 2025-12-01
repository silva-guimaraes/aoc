
# Joguei a toalha. Nunca consigo fazer Dijkstra.
# Toda a ideia desse código não é de minha autoria.

import heapq
input = open('16.txt').read().strip().split('\n')

start: tuple = (-1, -1)
for i, line in enumerate(input):
    for j, c in enumerate(line):
        if c == 'S':
            start = (i,j)

queue  = [(0, (start[0], start[1]), (0, 1), set())]
visited = {(start[0], start[1], 0, 1)}

best = None
rest = set()

while queue:
    (cost, (i, j), (di, dj), path) = heapq.heappop(queue)

    visited.add((i, j, di, dj))
    if input[i][j] == 'E':
        if best == None:
            best = cost
        if cost != best:
            break
        rest |= path

    new_states = [
        (cost+1,    (i+di, j+dj),   (di, dj), {(i, j)} | path),
        (cost+1000, (i, j),         (-dj, di), path),
        (cost+1000, (i, j),         (dj, -di), path)
    ]
    for new_state in new_states:
        _, (ni, nj), (ndi, ndj), _ = new_state
        if input[ni][nj] == '#':
            continue
        if (ni, nj, ndi, ndj) in visited:
            continue

        heapq.heappush(queue, new_state)

print(best)
print(len(rest)+1)
