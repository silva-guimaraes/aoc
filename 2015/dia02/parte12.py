
input = open('input.txt').read().strip().split('\n')

print(sum([y[0]*y[1] + 2*y[0]*y[1] + 2*y[0]*y[2] + 2*y[1]*y[2] for y in
           [sorted(map(int, x.split('x'))) for x in input]]))

print(sum([(y[0]+y[1])*2 + y[0]*y[1]*y[2] for y in [sorted(map(
    int, x.split('x'))) for x in input]]))
