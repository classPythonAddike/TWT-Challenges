def check(base, num):
    n = 1
    while n < num: n *= base
    return not(num * base > 2 * n) and num * base + 1 == n + num

for i in range(int(input())):
    num = int(input())
    for i in range(2, int(num ** 0.3) * 15):
        if check(i, num):
            print(i)
            break
    else:
        print(num - 1)
