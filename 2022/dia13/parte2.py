
right_order = True
undefined = "undefined"
wrong_order = False

# python pelo visto só consegue ordenar uma lista se os elementos tiverem algum valor numérico
# e se for possivel comparar esses numeros com maior que ou menor que
def bubble_sort(lst, cmp_func):
    for i in range(len(lst) - 1):
        for j in range(len(lst) - 1 - i):
            if not cmp_func(lst[j], lst[j+1]):
                lst[j], lst[j+1] = lst[j+1], lst[j]
    return lst
# em lisp isso seria um simples (sort #'compare-packets inputs) 

def compare_packets(a, b):
    for i in range(min(len(a), len(b))):

        left = a[i]
        right = b[i]

        if isinstance(left, list) and isinstance(right, list):
            ret = compare_packets(left, right)
            if ret == undefined:
                continue
            else:
                return ret

        elif isinstance(left, list) and isinstance(right, int):
            ret =  compare_packets(left, [right])
            if ret == undefined:
                continue
            else:
                return ret

        elif isinstance(left, int) and isinstance(right, list):
            ret =  compare_packets([left], right)
            if ret == undefined:
                continue
            else:
                return ret

        elif left == right:
            continue
        elif left < right:
            return right_order
        else:
            return wrong_order

    if len(a) < len(b):
        return right_order
    elif len(a) == len(b):
        return undefined
    else:
        return wrong_order


inputs = []

with open('./packets.txt') as file:
    for line in file:
        inputs.append(line.rstrip())

inputs = list(filter(lambda x: (x != ""), inputs))

inputs = list(map(lambda x: (eval(x)), inputs))

inputs.append([[6]])
inputs.append([[2]])

inputs = bubble_sort(inputs, compare_packets)

print((inputs.index([[6]]) + 1) * (inputs.index([[2]]) + 1))


