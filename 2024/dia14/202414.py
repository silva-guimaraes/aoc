
import re

def part1():
    r = re.compile(r'-?\d+')
    quadrants = [0,0,0,0]
    for i in open('14.txt').read().strip().split('\n'):
        x,y,vx,vy = list(map(int, r.findall(i)))
        x = (x + vx*100) % 101
        y = (y + vy*100) % 103
        if x == 50 or y == 51:
            continue
        quadrants[int(x > 50) + int(y > 51) * 2] += 1
    print(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])

# Essa foi a segunda parte mais burra que eu já fiz.
# Pra resolver esse problema do mesmo jeito que eu fiz, remova do comentário o código
# a baixo e preste bastante atenção na sua tela por 10 minutos ininterruptos.
# Pressione 'P' para o modo automático.
# Favor não piscar sem antes pausar.
# 'H', 'L', retornam e avançam um frame, respectivamente.
# Isso requer essa biblioteca: https://github.com/kitao/pyxel
# Boa sorte!
def part2():
    # import pyxel
    # tall = 103
    # wide = 101
    # input = open('14.txt').read().strip().split('\n')
    #
    # r = re.compile(r'-?\d+')
    # robots: List[List[int]] = []
    # for i in input:
    #     robots.append(list(map(int, r.findall(i))))
    #
    # at = 10403 # too high
    # at = 6911
    #
    # def robots_at(i: int):
    #     return [((x+vx*i) % wide, (y+vy*i) % tall) for [x, y, vx, vy] in robots]
    #
    # pyxel.init(wide, tall)
    # auto = False
    #
    # def update():
    #     nonlocal at, auto
    #     if pyxel.btnp(pyxel.KEY_Q):
    #         pyxel.quit()
    #     if pyxel.btnp(pyxel.KEY_L):
    #         at += 1
    #         print(at)
    #         print(at)
    #     if pyxel.btnp(pyxel.KEY_H):
    #         at -= 1
    #         print(at)
    #     if pyxel.btnp(pyxel.KEY_P):
    #         auto = not auto
    #         print('auto:', auto)
    #     if auto:
    #         at += 1
    #         print(at)
    #
    # def draw():
    #     pyxel.cls(0)
    #     for (x, y) in robots_at(at):
    #         pyxel.rect(x, y, 1, 1, 7)
    #
    # pyxel.run(update, draw)

    print(8053) # É o que deu pra fazer...
    # Enquanto eu desenvolvia esse método, percebi que existem certas 'faixas' que aparecem
    # uma vez a cada N frames.
    # São duas dessas faixas, uma aparece a cada 101 frames e a outra a cada 103 frames. Curiosamente a largura
    # e altura do grid. Curiosamente ambos números primos.
    # Existe uma intersecção entre as duas faixas que é onde o simbolo está escondido.
    # Talvez haja uma forma de calcular onde essa intersecção ocorre, minimo múltiplo comum foi a primeira
    # coisa que me veio em mente, existem alguns poréns que me deixaram perplexos e por isso 
    # resolvi fazer do modo mais burro.
    # Esse dia evidentemente foi uma manha do Eric pra deixar todos que usam IA de calças arriadas.


# Ok, dei uma pesquisada. aparentemente os robos não se sobrepõem quando o símbolo aparece
# def part2_2():
#     input = open('14.txt').read().strip().split('\n')
#     tall = 103
#     wide = 101
#
#     r = re.compile(r'-?\d+')
#     robots: list[list[int]] = []
#     for i in input:
#         robots.append(list(map(int, r.findall(i))))
#     for i in range(tall*wide):
#         s = set()
#         for [x, y, vx, vy] in robots:
#             s.add(((x+vx*i) % wide, (y+vy*i) % tall))
#         if len(s) == len(robots):
#             print(i)
# NÃO FUNCIONA
# ESPECIFICAMENTE COM O MEU INPUT NÃO FUNCIONA
# O QUÃO AZARADO EU CONSIGO SER
# AAAAA


part1()
part2()
# part2_2()
