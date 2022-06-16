K = 10**9 + 7

buf = ""

for i in range(int(input())):
    l = input().split()

    rem = lambda a: int(a) % K if int(a) >= 0 else -(abs(int(a)) % K)

    ops = {
        "*": lambda a, b: int(a) * int(b),
        "-": lambda a, b: int(a) - int(b),
        "+": lambda a, b: int(a) + int(b)
    }

    for op in ["*", "-", "+"]:
        while op in l:
            pos = l.index(op)
            l[pos] = ops[op](l[pos - 1], l[pos + 1])

            l.pop(pos - 1)
            l.pop(pos)

    buf += str(rem(l[0])) + "\n"

print(buf, end="")
