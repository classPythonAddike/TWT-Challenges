import sys

sys.setrecursionlimit(2000000)

MOD = 10 ** 9 + 7

trib = {0: 0, 1: 1, 2: 1}
trib_s = {1: 1}

def nth_trib(n):
    if (r := trib.get(n)) is None:
        r = (nth_trib(n - 1) + nth_trib(n - 2) + nth_trib(n - 3)) % MOD
        trib[n] = r
    return r

def nth_trib_s(n):
    if (r := trib_s.get(n)) is None:
        r = ((nth_trib(n) ** 3) % MOD + nth_trib_s(n - 1)) % MOD
        trib_s[n] = r
    return r


for i in range(int(input())):
    print(nth_trib_s(int(input())))
