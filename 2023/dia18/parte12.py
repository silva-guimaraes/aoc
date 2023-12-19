
from itertools import accumulate, pairwise

input = open('input.txt').read().strip()

print(int(sum([w[0][0][0]*w[1][0][1] - w[1][0][0]*w[0][0][1] + w[1][1] for w in pairwise(accumulate([((0, 0), 0)] + [({'R': (0, z[0]), 'D': (-z[0], 0), 'L': (0, -z[0]), 'U': (z[0], 0)}[z[1]], z[0]) for z in [(int(y[1]), y[0]) for y in [x.split(' ') for x in input.split('\n')]]], lambda a, b: ((a[0][0] + b[0][0], a[0][1] + b[0][1]), b[1])))]) / 2) + 1)

print(int(sum([w[0][0][0]*w[1][0][1] - w[1][0][0]*w[0][0][1] + w[1][1] for w in pairwise(accumulate([((0, 0), 0)] + [({'0': (0, z[0]), '1': (-z[0], 0), '2': (0, -z[0]), '3': (z[0], 0)}[z[1]], z[0]) for z in [(int(y[:-1], base=16), y[-1]) for y in [x.split(' ')[-1][2:-1] for x in input.split('\n')]]], lambda a, b: ((a[0][0] + b[0][0], a[0][1] + b[0][1]), b[1])))]) / 2) + 1)
