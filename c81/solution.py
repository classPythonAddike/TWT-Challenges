import math

def num_even_permutations(num: str) -> int:
    if len(num) == 1:
        return (int(num) + 1) % 2

    remainders = list(map(lambda x: int(x) % 2, num))

    even_digits = remainders.count(0)
    odd_digits = remainders.count(1)
    zeroes = num.count("0")

    denominator = 1

    for i in "0123456789":
        denominator *= math.factorial(num.count(i))

    fact = math.factorial(len(num) - 2)

    numerator = (odd_digits * even_digits * fact) + ((even_digits - zeroes) * (even_digits - 1) * fact)
    return int(numerator // denominator)

for i in range(int(input())):
    _sum = 0
    current_number = ""

    num = input()

    for i in num:
        current_number += i
        _sum += num_even_permutations(current_number)

    print(_sum)
