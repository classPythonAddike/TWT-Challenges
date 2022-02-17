# Just testing something :cheemsadidas:
# |````````|

for i in range(int(input())):

    roman_numerals = [
        ["I", 1],
        ["V", 5],
        ["X", 10],
        ["L", 50],
        ["C", 100],
        ["D", 500],
        ["M", 1000],
        ["IV", 4],
        ["IX", 9],
        ["XL", 40],
        ["XC", 90],
        ["CD", 400],
        ["CM", 900]
    ][::-1]

    rom_num = input()
    _sum = 0

    for rom, num in roman_numerals:
        _sum += rom_num.count(rom) * num
        rom_num = rom_num.replace(rom, "|")

    print(_sum)
