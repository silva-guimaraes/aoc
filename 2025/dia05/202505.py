
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
        if a[0] > b[0]:
            a, b = b, a
        x1, y1 = a
        x2, y2 = b
        assert x1 < y1
        assert x2 < y2
        if x1 < x2 and y1 > y2:
            rang2.remove(b)
            break
        elif x1 < y2 and y1 > x2:
            rang2.remove(a)
            rang2.remove(b)
            rang2.append((x1, y2))
            break
    if len(rang2) == s:
        break

pt2 += sum([b - a + 1 for a, b in rang2])

assert(all([(b - a + 1) >= 0 for a, b in rang2]))
print(pt2)


