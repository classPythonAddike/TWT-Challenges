for i in range(int(input())):
    n, m = list(map(int, input().split()))
    arr = list(map(int, input().split()))
    consec_sum = [0] * (n + 1)
    consec_sum[0] = 0

    for i in range(n):
        consec_sum[i + 1] = arr[i] ^ consec_sum[i]

    for j in range(m):
        a, b = list(map(int, input().split()))
        print(consec_sum[b] ^ consec_sum[a - 1])
