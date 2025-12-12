from functools import cache

input = open('11.txt').read().strip().split('\n')

map = {f: to.split() for f, to in map(lambda x: x.split(': '), input)}

@cache
def recur1(at):
    return 1 if at == 'out' else sum(recur1(i) for i in map[at])

@cache
def recur2(at, dac = False, fft = False):
    if at == 'out' and dac and fft:
        return 1 
    elif at == 'out':
        return 0
    else:
        return sum(recur2(i, dac or at == 'dac', fft or at == 'fft') for i in map[at])

print(recur1('you'))

print(recur2('svr'))
