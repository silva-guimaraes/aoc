import functools

patterns, towels = inpurt = open('19.txt', 'r').read().strip().split('\n\n')
patterns = list(map(lambda x: x.strip(), patterns.split(',')))

towels = towels.split()

@functools.cache
def recur(l):
    if not l:
        return 1
    s = 0
    for p in patterns:
        if p == l[:len(p)]:
            s += recur(l[len(p):])
    return s

print(sum(recur(t) > 0 for t in towels))
print(sum(recur(t) for t in towels))
