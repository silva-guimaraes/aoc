
from typing import List
import pyxel
from itertools import pairwise

input = open('../02.txt', 'r').read().strip().split('\n')
input = [list(map(int, i.split())) for i in input]

results: List = []

def safe(i: List[int]) -> tuple[bool, list[bool], list[bool], list[bool]]:
    a = list(map(lambda x: x[1] - x[0], pairwise(i)))
    diff = list(map(lambda x: 0 < abs(x) < 4, a))
    asc = list(map(lambda x: x > 0, a))
    dsc = list(map(lambda x: x < 0, a))
    if  (all(asc) or all(dsc)) and all(diff):
        return True, diff, asc, dsc
    return False, diff, asc, dsc

def part1():
    global results
    sum = 0
    for i in input:
        s, diff, asc, dsc = safe(i)
        if s: sum += 1

        result = { 'safe': s, 'input': i, 'diff': diff, 'asc': asc, 'dsc': dsc}
        results.append(result)
    print(sum)


def part2():
    sum = 0
    for a in input:
        if safe(a):
            sum += 1
            continue
        for i in range(len(a)):
            b = a[:]
            del b[i]
            if safe(b):
                sum += 1
                break
    print(sum)


part1()
# deveria ter usado o output da segunda parte!!!
# part2()

pyxel.init(256, 256)

# todas as fontes generozamente doadas por este repositório:
# https://github.com/olikraus/u8g2/wiki/fntgrpu8g#emoticons21
five_by_eight = pyxel.Font('./5x8.bdf')
freedom25 = pyxel.Font('./freedoomr25n.bdf')
freedom25_width = 18

def update():
    pyxel.screencast
    if pyxel.btnp(pyxel.KEY_Q):
        pyxel.quit()


speed_up = 1.0
valid = []
total_safe = sum(map(lambda x: x['safe'], results))
total_unsafe = len(input) - total_safe
red = 8
green = 11
end_credits_countdown = 0
end_credits = 32
margin = 3
gray = 13

def draw():
    global end_credits_countdown, speed_up

    pyxel.cls(0)

    if end_credits_countdown > end_credits:
        pyxel.text(margin, margin, f'safe: {total_safe}', gray, five_by_eight)
        pyxel.text(margin, margin+8+1, f'unsafe: {total_unsafe}', gray, five_by_eight)
        return


    for i, v in enumerate(valid):

        padded = ''.join([str(j).rjust(3) for j in v['input']])
        y = (4+4)*i

        if y > 256:
            break

        c = green if v['safe'] else red

        pyxel.text(256-(len(padded)*5)-margin, y, padded, c, five_by_eight)

    for i, r in enumerate(results):

        padded = ''.join([str(j).ljust(3) for j in r['input']])
        y = (4+4)*i

        if y > 256:
            break

        pyxel.text(margin, y, padded, gray, five_by_eight)

    # nao funciona do jeito que eu gostaria
    for _ in range(int(speed_up)):
        if len(results) > 0:
            valid.insert(0, results.pop(0))
        else:
            valid.insert(0, { 'safe': False, 'input': [], 'diff': [], 'asc': [], 'dsc': []})
            end_credits_countdown += 1

    speed_up *= 1.001


pyxel.run(update, draw)

