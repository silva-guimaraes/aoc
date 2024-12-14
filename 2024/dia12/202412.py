

from typing import Dict, Set, Tuple, List


input = open('12.txt', 'r').read().strip().split()


def inside(ij):
    i, j = ij
    if i < 0 or i >= len(input) or j < 0 or j >= len(input[0]):
        return False
    return True

def is_perimeter(c: str):
    def _get(ij: Tuple[int, int]):
        if not inside(ij):
            return True
        i, j = ij
        x = input[i][j]
        if x != c:
            return True
        return False
    return _get


def part1():
    visited: Set[Tuple[int,int]] = set()
    score = 0
    for i1, line in enumerate(input):
        for j1, c in enumerate(line):
            if (i1,j1) in visited:
                continue
            queue: List[Tuple[int,int]] = [(i1,j1)]
            perimeter = 0
            area = 0
            while len(queue) > 0:
                i, j = ij = queue.pop(0)
                if not inside(ij):
                    continue
                if input[i][j] != c:
                    continue
                if ij in visited:
                    continue
                visited.add(ij)
                directions = [(i+1,j), (i,j+1), (i-1,j), (i,j-1)]
                perimeter += sum(map( is_perimeter(c), directions))
                area += 1
                list(map(lambda x: queue.append(x), directions))
            score += area * perimeter
    print(score)


# é um flood fill em cima de um flood fill, pra cada uma das 4 direções.
# o primeiro flood fill salva os perímetros que cada campo faz.
# é lerdo, mas ao menos funciona.......
def part2():
    shadow: List[List[Tuple[List[bool], Tuple[int,int]]]] = [[([], (-1,-1)) for _ in i]for i in input]
    visited: Set[Tuple[int,int]] = set()
    areas: Dict[Tuple[int, int], int] = {}
    for i1, line in enumerate(input):
        for j1, c in enumerate(line):
            if (i1,j1) in visited:
                continue
            queue: List[Tuple[int,int]] = [(i1,j1)]
            perimeter = 0
            area = 0
            while len(queue) > 0:
                i, j = ij = queue.pop(0)
                if not inside(ij):
                    continue
                if input[i][j] != c:
                    continue
                if ij in visited:
                    continue
                visited.add(ij)
                directions = [(i+1,j), (i,j+1), (i-1,j), (i,j-1)]
                perimeter = list(map( is_perimeter(c), directions))
                shadow[i][j] = (perimeter, (i1, j1))
                area += 1
                list(map(queue.append, directions))
            areas[(i1,j1)] = area
    perimeters  = {k: 0 for k in areas.keys()}
    for (i1, j1) in areas.keys():
        for direction in range(4):
            visited.clear()
            for i2, line in enumerate(input):
                for j2, c in enumerate(line):
                    _, origin = shadow[i2][j2]
                    if origin != (i1, j1):
                        continue
                    queue: List[Tuple[int,int]] = [(i2, j2)]
                    any = False
                    while len(queue) > 0:
                        i, j = ij = queue.pop(0)
                        if not inside(ij):
                            continue
                        if input[i][j] != c:
                            continue
                        if ij in visited:
                            continue
                        d, _ = shadow[i][j]
                        if not d[direction]:
                            continue
                        visited.add(ij)
                        any = True
                        directions = [(i+1,j), (i,j+1), (i-1,j), (i,j-1)]
                        list(map(queue.append, directions))
                    if any:
                        perimeters[(i1,j1)] += 1
    score = 0
    for k in perimeters.keys():
        score += perimeters[k] * areas[k]
    print(score)




part1()
part2()
