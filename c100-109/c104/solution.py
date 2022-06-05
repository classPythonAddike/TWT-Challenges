# Kadane's Algo

for i in range(int(input())):
    input()
    A = list(map(int, input().split()))

    a = min(A)
    b = 0

    for i in A:
        b += i
        if a < b: a = b
        if b < 0: b = 0
    print(max(a, 0))
