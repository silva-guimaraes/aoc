
from functools import cache
from itertools import combinations
_wires, _connections = open('24.txt').read().strip().split('\n\n')

connections: dict[str, int | tuple[str, str, str]] = dict()
for i in _connections.split('\n'):
    a, op, b, _, out = i.split()
    connections[out] = (op, a, b)

for i in _wires.split('\n'):
    name, state = i.split(': ')
    connections[name] = int(state)

def recur(query: str) -> int:
    x = connections[query]
    match x:
        case int(): 
            return x
        case ('AND', a, b):
            return int(recur(a) and recur(b))
        case ('OR', a, b):
            return int(recur(a) or recur(b))
        case ('XOR', a, b):
            return int(recur(a) != recur(b))
        case _:
            raise

@cache
def dependency(query: str) -> set[str]:
    match connections[query]:
        case int():
            return {query}
        case (_, a, b):
            return {query, *dependency(a), *dependency(b)}

def can_pair(query: str, visited: set[str]):
    x = dependency(query)
    for i in set(connections.keys()).difference(x):
        if i[0] in 'xy':            continue
        if i in visited:            continue
        if query in dependency(i):  continue
        yield (query, i)

visited = set()
pairs = []
for i in connections.keys():
    pairs += can_pair(i, visited)
    visited.add(i)

sum = 0
for i in combinations(pairs, 4):
    sum += 1
print(sum)

# print(list(foo('z12')))
# print(int(str().join(map(lambda x: str(recur(x)),
#                          sorted(filter(lambda x: x[0] == 'z', connections.keys()),
#                                 reverse=True))), base=2))
