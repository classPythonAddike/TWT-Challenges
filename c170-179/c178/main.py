import math

for i in range(int(input())):
    n = input()
    l = len(n)

    p = 1
    letters = dict()

    for i in n[:l//2]:
        letters[i] = letters.get(i, 0) + 1
    
    tot = 0
    for r in letters.values():
        comb = 1
        for i in range(1, 1 + r):
            tot += 1
            comb = (comb * tot) // i
        p = (p * comb) % (10**9 + 7)

    print(p)
        
