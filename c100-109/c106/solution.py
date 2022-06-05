import math

buf = ""

for i in range(int(input())):
    s, n = map(int, input().split())
    
    if n < 0 or n > (s * (s + 1)) // 2:
        buf += "BAD INPUT\n"
        continue
    if n == 0:
        buf += "2 1\n"
        continue

    D = ((2 * s - 1) ** 2) - 8 * n + 8 * s
    i = ((2 * s - 1) - D ** 0.5) / 2
    if i % 1 == 0:
        i = int(i) - 1
    row = math.floor(i) + 1

    c = ((row + 1) * (2 * s - row)) // 2
    column = s - (c - n)
    buf += f"{row + 1} {column}\n" # row is 0 indexed, so need to add 1

print(buf)
