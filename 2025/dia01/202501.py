
input = open('01.txt', 'r').read().strip().split()

current1, current2 = 50, 50
pt1, pt2 = 0, 0

for i in input:
    direction = 1 if i[0] == 'R' else -1
    number = int(i [1:])
    current1 = (current1 + direction * number) % 100
    if current1 == 0:
        pt1 += 1
    for n in range(number):
        current2 = (current2 + direction * 1) % 100
        if current2 == 0:
            pt2 += 1

print(pt1)
print(pt2)

