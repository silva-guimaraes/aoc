
input = open('15.txt', 'r').read().strip()

grid, moves = input.split('\n\n')
grid = grid.split('\n')
moves = ''.join(moves.split('\n'))
directions = {
    'v': (1, 0),
    '>': (0, 1),
    '<': (0, -1),
    '^': (-1, 0),
}

def part1():
    walls = set()
    boxes = set()
    grobot: tuple[int,int] = (-1,-1)
    robot: tuple[int,int] = grobot
    for i, line in enumerate(grid):
        for j, c in enumerate(line):
            if c == '#':
                walls.add((i,j))
            elif c == 'O':
                boxes.add((i,j))
            elif c == '@':
                robot = (i,j)

    for m in moves:
        mi, mj = directions[m]
        ri, rj = robot
        ni, nj = new = (ri + mi, rj + mj)
        if new in boxes:
            t = new
            while t in boxes:
                t = (t[0]+mi, t[1]+mj)
            if not t in walls:
                boxes.add(t)
                boxes.remove(new)
                robot = (ni, nj)
        elif new in walls:
            continue
        else:
            robot = (ni, nj)

    print(sum(i*100 + j for (i, j) in boxes))


type Pos = tuple[int,int]
type DoubleBox = tuple[Pos, Pos]
def part2():
    class MySet:
        def __init__(self):
            self.itens: set[DoubleBox] = set()
        def add(self, p1: Pos, p2: Pos):
            self.itens.add((p1, p2))
        def remove(self, item):
            self.itens.remove(item)
        def get(self, x: tuple[int,int]):
            for pos in self.itens:
                (p1, p2) = pos
                if x == p1 or x == p2:
                    return pos
            raise 
        def __contains__(self, x: Pos):
            for (p1, p2) in self.itens:
                if x == p1 or x == p2:
                    return True
            return False
        def __iter__(self):
            return self.itens.__iter__()

    walls = MySet()
    boxes = MySet()
    robot: tuple[int,int] = (-1,-1)
    for i, line in enumerate(grid):
        for j, c in enumerate(line):
            if c == '#':
                walls.add((i,j*2), (i,j*2+1))
            elif c == 'O':
                boxes.add((i,j*2), (i,j*2+1))
            elif c == '@':
                robot = (i,j*2)

    for _, m in enumerate(moves):
        mi, mj = directions[m]
        ri, rj = robot
        new = (ri + mi, rj + mj)
        if new in boxes:
            q1, q2 = boxes.get(new)
            queue = [q1, q2]
            unit = MySet()
            canMove = True
            while len(queue) > 0:
                p = queue.pop(0)
                if p in unit:
                    continue
                elif p in boxes:
                    ((p1i, p1j), (p2i, p2j)) = db = boxes.get(p)
                    unit.add(db[0], db[1])
                    queue.append((p1i+mi, p1j+mj))
                    queue.append((p2i+mi, p2j+mj))
                elif p in walls:
                    canMove = False
                    break
            if canMove:
                for db in unit:
                    boxes.remove(db)
                for db in unit:
                    ((p1i, p1j), (p2i, p2j)) = db 
                    boxes.add((p1i+mi, p1j+mj), (p2i+mi, p2j+mj))
                robot = new
                    
        elif new in walls:
            continue
        else:
            robot = new
    print(sum(i*100 + j for ((i, j), _) in boxes))

part1()
part2()
