
from heapq import heappop, heappush


# a única primeira parte desse ano que eu não consegui resolver logo no dia

# ralei pra entender que um campo poderia ser visitado mais de uma vez dado
# que a direção e o numero de passos consecutivos fossem todos diferentes

input = open('input.txt').read().strip().split('\n')

grid = [[int(b) for b in a] for a in input]

start, end = (0, 0), (len(grid[0])-1, len(grid)-1)
heap = [(0, start, 0, (1, 0), (0, 0))]
visited = set()

while len(heap) > 0:
    heat, pos, straights, direction, past = heappop(heap)
    x, y = pos
    (dx, dy) = direction

    visit = (pos, straights, direction)
    if visit in visited:
        continue
    visited.add(visit)

    if pos == end:
        print(heat)
        break

    for new_dir in [(-1, 0), (1, 0), (0, -1), (0, 1)]:

        if direction == (-new_dir[0], -new_dir[1]):
            continue

        if straights == 3 and direction == new_dir:
            continue

        new = (pos[0] + new_dir[0], pos[1] + new_dir[1])

        if not 0 <= new[0] < len(grid[0]) or not 0 <= new[1] < len(grid):
            continue

        s = straights+1 if direction == new_dir else 1

        heappush(heap, (heat + grid[new[1]][new[0]], new, s, new_dir, pos))
