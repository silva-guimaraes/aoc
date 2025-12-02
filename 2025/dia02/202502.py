
pt1, pt2 = 0, 0

for i in open('02.txt', 'r').read().strip().split(','):
    x, y = i.split('-')
    for j in range(int(x), int(y)+1):
        s = str(j)
        l = len(s) // 2
        if len(s) % 2 == 0 and s[:l] == s[l:]:
            pt1 += j
        if any([len(s) % i == 0 and (s[:i])*(len(s) // i) == s for i in range(1, len(s))]):
            pt2 += j

print(pt1)
print(pt2)
