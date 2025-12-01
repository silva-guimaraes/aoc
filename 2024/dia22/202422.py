
from itertools import pairwise

input = list(map(int, open('22.txt').read().strip().split()))

prune = lambda x: x & 0xffffff

def secret(n, r):
    nn = []
    for _ in range(r):
        a = prune((n << 6)^n)
        b = prune((a >> 5) ^ a)
        n = prune((b << 11) ^ b)
        nn.append(n % 10)
    return n, nn, list(map(lambda x: x[1] - x[0], pairwise(nn)))

def window(x, n):
    for i, _ in enumerate(x):
        if i+n > len(x): break
        yield tuple(x[i:i+n])

buyers = [secret(i, 2000) for i in input]

print(sum(price for (price, _, _) in buyers))

buyers_sequences = []
sequences_sell_sums: dict[tuple[int,...], int] = {}

for (_, bananas, diff) in buyers:
    window_diff = list(window(diff, 4))
    diff_banana: dict[tuple[int,...], int] = {}

    for i, d in enumerate(window_diff):
        b = bananas[i+4]
        if not d in diff_banana:
            diff_banana[d] = b

    for k in diff_banana.keys():
        if k in sequences_sell_sums:
            sequences_sell_sums[k] += diff_banana[k]
        else:
            sequences_sell_sums[k] = diff_banana[k]

print(max([sequences_sell_sums[v] for v in sequences_sell_sums.keys()]))
