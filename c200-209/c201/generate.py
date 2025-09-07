for n in range(256):
    print(int('{:08b}'.format(n)[::-1], 2), end=",")
