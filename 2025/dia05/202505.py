
import itertools

rang1, input = open('05.txt').read().split('\n\n')
rang2: list[tuple[int, int]] = []
for i in rang1.split():
    a, b = i.split('-')
    rang2.append((int(a), int(b)))
input = list(map(int, input.split()))


pt1 = 0
for i in input:
    for a, b in rang2:
        if a <= i <= b:
            pt1 += 1
            break

print(pt1)

pt2 = 0
while True:
    s = len(rang2)
    for a, b in itertools.combinations(rang2, 2):
        ax, ay = a
        bx, by = b
        if a == b:
            rang2.remove(b)
        # A contem B
        elif ax <= bx and ay >= by:
            rang2.remove(b)
        # B contem A
        elif bx <= ax and by >= ay:
            rang2.remove(a)
        # A antes
        elif ax < bx and ay >= bx:
            rang2.remove(a)
            rang2.remove(b)
            rang2.append((ax, by))
        elif ax < bx and ay < bx:
            continue
        # B antes
        elif bx < ax and by >= ax:
            rang2.remove(a)
            rang2.remove(b)
            rang2.append((bx, ay))
        elif bx < ax and by < ax:
            continue
        else:
            continue
        break
    if len(rang2) == s:
        break

pt2 += sum([b - a + 1 for a, b in rang2])

print(pt2)
