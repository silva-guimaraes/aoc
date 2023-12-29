
from heapq import heappop, heappush

input = open('input.txt').read().strip().split('\n')

grid = [[int(b) for b in a] for a in input]

start, end = (0, 0), (len(grid[0])-1, len(grid)-1)
heap = [(0, start, 0, (1, 0)), (0, start, 0, (0, 1))]
visited = set()

while len(heap) > 0:
    heat, pos, straights, direction = heappop(heap)

    visit = (pos, straights, direction)
    if visit in visited:
        continue
    visited.add(visit)

    if pos == end and straights >= 4:
        print(heat)
        break

    for new_dir in [(-1, 0), (1, 0), (0, -1), (0, 1)]:

        if direction == (-new_dir[0], -new_dir[1]):
            continue

        new = (pos[0] + new_dir[0], pos[1] + new_dir[1])

        if not 0 <= new[0] < len(grid[0]) or not 0 <= new[1] < len(grid):
            continue

        s = straights+1 if direction == new_dir else 1

        if new_dir != direction and straights < 4:
            continue

        if new_dir == direction and straights == 10:
            continue

        heappush(heap, (heat + grid[new[1]][new[0]], new, s, new_dir))
