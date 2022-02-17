# Finished in 7 minutes :cheemsadidas:
for i in range(int(input())):
    char = input()[0]
    moves = input().split()

    if char == "X":
        x_moves = moves[0::2]
        o_moves = moves[1::2]
    else:
        x_moves = moves[1::2]
        o_moves = moves[0::2]

    x_moves = set(x_moves)
    o_moves = set(o_moves)

    wins = ["123", "456", "789", "147", "258", "369", "159", "357"]

    for m in wins:
        if len(set(m).intersection(x_moves)) == 3:
            print("X")
            break
        elif len(set(m).intersection(o_moves)) == 3:
            print("O")
            break
    else:
        print("Tie")
