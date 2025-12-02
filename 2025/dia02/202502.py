
input = open('02.txt', 'r').read().strip().split(',')

pt1, pt2 = 0, 0

def invalid1(s):
    l = int(len(s)/2)
    return s[:l] == s[l:]

def invalid2(s: str):
    for i in range(1, len(s)):
        if len(s) % i != 0:
            continue
        m = int(len(s) / i)
        x = s[:i]
        if x*m == s:
            return True
    return False

for i in input:
    x, y = i.split('-')
    for j in range(int(x), int(y)+1):
        s = str(j)
        if invalid1(s):
            pt1 += j
        if invalid2(s):
            pt2 += j

print(pt1)
print(pt2)
