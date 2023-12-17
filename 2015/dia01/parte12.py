import itertools

input = open('input.txt').read().strip()

print(sum(1 if x == "(" else -1 for x in input))
print(
    list(filter(lambda x: x[1] == -1, enumerate(itertools.accumulate(
        map(lambda x: 1 if x == "(" else -1, input)))))[0][0] + 1)
