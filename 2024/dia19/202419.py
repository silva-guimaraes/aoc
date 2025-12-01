from functools import cache

patterns, towels = open('19.txt').read().strip().split('\n\n')
patterns = patterns.split(', ')
towels = towels.split()

@cache
def recur(l):
    return 1 if not l else sum(recur(l[len(p):]) for p in patterns if p == l[:len(p)])

x = list(map(recur, towels))

print(sum(i > 0 for i in x))
print(sum(x))
