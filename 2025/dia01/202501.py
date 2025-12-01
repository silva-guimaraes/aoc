
input = open('01.txt', 'r').read().strip().split()

current = 50
zeros = 0 

for i in input:
    direction = i[0]
    number = int(i [1:])
    if direction == 'R':
        current += number
    elif direction == 'L':
        current -= number
    current %= 100
    if current == 0:
        zeros += 1

print(zeros)

current = 50
zeros = 0

for i in input:
    direction = i[0]
    number = int(i [1:])
    for n in range(number):
        if direction == 'R':
            current += 1
        elif direction == 'L':
            current -= 1
        current %= 100
        if current == 0:
            zeros += 1

print(zeros)
