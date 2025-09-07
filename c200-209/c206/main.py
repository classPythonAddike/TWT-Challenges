from bisect import bisect_left, bisect_right

for i in range(int(input())):
    N = int(input())
    sides = sorted(map(int, input().split()))

    print(sides)
    count = 0

    info = [(N - i - 1, i) for i, j in enumerate(sides)]

    for i in range(N):
        j = i + 1
        while j < N:
            s = sides[i] + sides[j]
            d = abs(sides[i] - sides[j])

            s_i = bisect_left(sides, s)
            d_i = max(bisect_right(sides, d), j)

            print(sides[i], sides[j], sides[d_i : s_i], d_i <= j < s_i)
            count += s_i - d_i - (d_i <= j < s_i)
            j += 1

    print(count)
