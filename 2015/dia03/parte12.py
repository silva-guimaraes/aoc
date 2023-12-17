from itertools import accumulate, islice
from functools import reduce
input = open('input.txt').read().strip()


# essa função só faz parte do itertools na versão 3.12 e eu só tenho
# acesso a versão 3.11
def batched(iterable, n):
    # batched('ABCDEFG', 3) --> ABC DEF G
    if n < 1:
        raise ValueError('n must be at least one')
    it = iter(iterable)
    while batch := tuple(islice(it, n)):
        yield batch


print(len(set(accumulate(
    [(0, 0)] + [{'^': (1, 0), 'v': (-1, 0), '<': (0, -1), '>': (0, 1)}[x]
                for x in input], func=lambda x, y: (x[0] + y[0], x[1] + y[1])
    ))))

print(len(set(reduce(
    lambda x, y: x + [y[0], y[1]],
    accumulate(batched(
        [{'^': (1, 0), 'v': (-1, 0), '<': (0, -1), '>': (0, 1)}[x]
         for x in input], 2),
               func=lambda x, y: ((x[0][0] + y[0][0], x[0][1] + y[0][1]),
                                  (x[1][0] + y[1][0], x[1][1] + y[1][1])),
               initial=((0, 0), (0, 0))), []))))
