
import functools

from typing import Dict, List, Tuple


_rules, _updates = open('./05.txt', 'r').read().split('\n\n')
rules: List[Tuple[int, int]] = []
updates: List[List[int]] = []

for i in _rules.strip().split('\n'):
    a, b = i.split('|') 
    rules.append((int(a), int(b)))

for i in _updates.strip().split('\n'):
    updates.append(list(map(int, i.split(','))))

def middle_page(x):
    return x[len(x) // 2]

def part1():
    correct = []
    for update in updates:
        order = {u: i for i, u in enumerate(update)}
        
        in_order = True
        for before, after in rules:
            if not before in update or not after in update:
                continue

            if order[before] > order[after]:
                in_order = False
                break
        if in_order:
            correct.append(update)
    print(sum(map(middle_page, correct)))


def part2():
    incorrect = []
    before_after: Dict[int, List[int]] = {}
    for b, a in rules:
        if not b in before_after:
            before_after[b] = []

        before_after[b].append(a)

    print(before_after)

    for update in updates:
        order = {u: i for i, u in enumerate(update)}
        
        in_order = True
        for b, a in rules:
            if not b in update or not a in update:
                continue

            if order[b] > order[a]:
                in_order = False
                break
        if not in_order:
            incorrect.append(update)

    def cmp(a: int, b: int):
        if not a in before_after:
            return 0
        return -1 if b in before_after[a] else 1

    for i in incorrect:
        i.sort(key=functools.cmp_to_key(cmp))

    print(sum(map(middle_page, incorrect)))



part1()
part2()

