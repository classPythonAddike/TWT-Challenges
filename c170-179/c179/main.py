import math

def union(l1, l2):
    return [l1[i] or l2[i] for i in range(len(l1 + l2) // 2)]

for i in range(int(input())):
    n, m = list(map(int, input().split()))

    l = list(map(int, input().split()))
    n_happy, happy_boys = l[0], l[1:]
    boys = [i in happy_boys for i in range(n)]

    l = list(map(int, input().split()))
    m_happy, happy_girls = l[0], l[1:]
    girls = [i in happy_girls for i in range(m)]

    if n_happy == 0 and m_happy == 0:
        print("NO")
        continue

    if (gcd := math.gcd(n, m)) == 1 and (m_happy or n_happy):
        print("YES")
        continue

    step = [0] * gcd

    for i in range(0, m // gcd):
        step = union(step, girls[i * gcd: (i + 1) * gcd])
    for i in range(0, n // gcd):
        step = union(step, boys[i * gcd: (i + 1) * gcd])
    
    if all(step):
        print("YES")
    else:
        print("NO")
