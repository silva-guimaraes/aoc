
import math
from itertools import pairwise

input = list(map(int, open('22.txt').read().strip().split()))

def secret(n, r):
    nn = []
    for _ in range(r):
        a = ((n*64)^n) % 16777216
        b = (int(a / 32) ^ a) % 16777216 
        n = ((b * 2048) ^ b) % 16777216 
        nn.append(int(str(n)[-1]))
    return n, nn, list(map(lambda x: x[1] - x[0], pairwise(nn)))

def window(x, n):
    for i, _ in enumerate(x):
        if i+n > len(x): break
        yield tuple(x[i:i+n])

buyers = [secret(i, 2000) for i in input]

print(sum(price for (price, _, _) in buyers))

different_sequences = set()
buyers_sequences = []
for (_, bananas, diff) in buyers:
    window_diff = list(window(diff, 4))
    different_sequences |= set(window_diff)

    diff_banana = {}
    for i, d in enumerate(window_diff):
        b = bananas[i+4]
        if not d in diff_banana:
            diff_banana[d] = b

    buyers_sequences.append(diff_banana)

best_sequence = -math.inf
for diff in different_sequences:
    s = 0
    for buyer in buyers_sequences:
        if diff in buyer:
            s += buyer[diff]
    best_sequence = max(best_sequence, s)

print(best_sequence)
