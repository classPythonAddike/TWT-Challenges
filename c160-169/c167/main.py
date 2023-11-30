def isPalindrome(a: int) -> bool:
    return str(a) == str(a)[::-1]

N = 10 ** 5
palindromes = []

for i in range(N):
    if isPalindrome(i):
        palindromes.append(i)

print(len(palindromes))
exit()

for i in range(int(input())):
    a, b, m = map(int, input().split())
    c = 0

    for j in palindromes:
        if j >= a and j <= b:
            if isPalindrome(j // m):
                c += 1
        if j > b:
            break

    print(c)
