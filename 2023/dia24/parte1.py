from itertools import combinations
import re

x, y = 0, 1


# https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection
def formula(p1, p2, p3, p4):
    px = ((p1[x]*p2[y] - p1[y]*p2[x]) * (p3[x]-p4[x]) -
          (p1[x] - p2[x]) * (p3[x]*p4[y] - p3[y]*p4[x]))
    py = ((p1[x]*p2[y] - p1[y]*p2[x]) * (p3[y]-p4[y]) -
          (p1[y] - p2[y]) * (p3[x]*p4[y] - p3[y]*p4[x]))
    div = (p1[x] - p2[x]) * (p3[y] - p4[y]) - (p1[y] - p2[y]) * (p3[x]-p4[x])

    return (px / div, py / div) if div != 0 else None


# bleh
def past(p, v, inter):
    if v[0] > 0 and inter[0] < p[0]:
        return True
    if v[0] < 0 and inter[0] > p[0]:
        return True
    if v[1] > 0 and inter[1] < p[1]:
        return True
    if v[1] < 0 and inter[1] > p[1]:
        return True
    return False


input = open('input.txt').read().strip().split('\n')

nums = re.compile("-?\\d+")

scale = 1000000000000000000000000
# scale = 10

points = [[int(b) for b in nums.findall(a)] for a in input]


# atleast, atmost = 7, 27
atleast, atmost = 200000000000000, 400000000000000

count = 0
for a, b in combinations(points, 2):
    p1 = a[:3]
    p2 = [a[i] + a[i+3]*scale for i in range(3)]
    p3 = b[:3]
    p4 = [b[i] + b[i+3]*scale for i in range(3)]

    intersections = formula(p1, p2, p3, p4)

    if intersections is None:
        continue
    if past(p1, a[3:], intersections):
        continue
    if past(p3, b[3:], intersections):
        continue

    if (intersections[0] > atleast and intersections[0] < atmost and
        intersections[1] > atleast and intersections[1] < atmost):
        count += 1

print(count)
