

import sys
sys.setrecursionlimit(10000000)

input = open('23.txt').read().strip().split()
connections: dict[str, set[str]] = dict()
computers: set[str] = set()

for i in input:
    a, b = i.split('-')
    computers.add(a)
    computers.add(b)

    if not a in connections:
        connections[a] = set()
    connections[a].add(b)

    if not b in connections:
        connections[b] = set()
    connections[b].add(a)

visited = set()
triples = set()
sum = 0
for a in computers:
    this_visited = set()
    for b in connections[a]:
        if b in visited or b in this_visited:
            continue
        for c in (connections[b] & connections[a]):
            if c in visited or c in this_visited:
                continue
            triples.add((a,b,c))
            if a[0] == 't' or b[0] == 't' or c[0] == 't':
                sum += 1
        this_visited.add(b)
    visited.add(a)

print(sum)

# https://en.wikipedia.org/wiki/Bron%E2%80%93Kerbosch_algorithm
def bron_kerboschl(P: set[str], X: set = set(), R: set = set(),):
    if len(P) == 0 and len(X) == 0:
        return R
    m = set()
    for v in list(P):
        x = bron_kerboschl(P & connections[v], X & connections[v], R | {v})
        if len(x) > len(m):
            m = x
        P.remove(v)
        X.add(v)
    return m

print(','.join(sorted(bron_kerboschl(computers))))

