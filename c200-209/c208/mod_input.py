with open("input.txt", "r") as f:
    inp = f.read().strip() + '\n'


with open("input.txt", "w") as f:
    f.write("3000\n")
    f.write(inp * 1000)

