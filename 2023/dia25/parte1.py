
input = open('input.txt').read().strip().split('\n')

components = {j[0][:-1]: j[1:] for j in [i.split(' ') for i in input]}

# print(components)

[[components.setdefault(i, []).append(k) for i in v]
 for k, v in components.copy().items()]

components = {k: list(set(v)) for k, v in components.items()}

# usei o graphviz
# neato -Tpng  > output.png

components['klk'].remove('xgz')
components['xgz'].remove('klk')

components['vmq'].remove('cbl')
components['cbl'].remove('vmq')

components['bvz'].remove('nvf')
components['nvf'].remove('bvz')

visited = set()
queue = ['dtl']

while len(queue) > 0:
    pop = queue.pop(0)
    visited.add(pop)
    [queue.append(i) for i in components[pop] if i not in visited]

print(len(visited) * (len(components) - len(visited)))
