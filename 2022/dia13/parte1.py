
# em python por que não tive vontade de ajeitar os inputs em go e python ja tem o eval()

right_order = "right order"
undefined = "undefined"
wrong_order = "wrong order"

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

sum = 0

for i in range(0, len(inputs) - 1, 2):
    pair = compare_packets(inputs[i], inputs[i + 1])
    if pair == right_order:
        sum += round(i / 2 + 1)

print(sum)

