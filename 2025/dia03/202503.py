

input = open('03.txt').read().strip().split()

pt1 = 0

for bank in input:
    i = 0
    m = -1
    while i < len(bank)-1:
        j = i+1
        while j < len(bank):
            xi, xj = bank[i], bank[j]
            m = max(int(xi+xj), m)
            j += 1
        i += 1
    pt1 += m

print(pt1)


