
import math
from itertools import combinations

input = open('teste.txt').read().strip().split()
boxes = [tuple(map(int, i.split(','))) for i in input]

pairs = []
for a, b in combinations(range(len(boxes)), 2):
    xa, ya, za = boxes[a]
    xb, yb, zb = boxes[b]
    distance = math.sqrt(math.pow(xa - xb, 2) + math.pow(ya - yb, 2) + math.pow(za - zb, 2))
    pairs.append((distance, a, b))

pairs.sort()
connections = 10
print(pairs[:connections//2])
    

