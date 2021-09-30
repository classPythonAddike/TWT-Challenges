def sub_matrices(a):
    """Get all unique submatrices in `a`"""
    matrices = {}

    for i in range(len(a) - 1):
        for j in range(len(a[0]) - 1):
            p = a[i][j]
            q = a[i][j + 1]
            r = a[i + 1][j]
            s = a[i + 1][j + 1]

            # We need a unique identifier for each submatrix -
            matrices[f"{p}{q}{r}{s}"] = 0

    return len(matrices)


for i in [i := input] * int(i()):
    l, _ = map(int, i().split())

    # Split input into a list
    m = [i().split() for j in range(l)]

    print(sub_matrices(m))

# ---- This solution didn't pass due to a bug, a corrected version is provided below ----


def sub_matrices(a):
    """Get all unique submatrices in `a`"""
    matrices = {}

    for i in range(len(a) - 1):
        for j in range(len(a[0]) - 1):
            p = a[i][j]
            q = a[i][j + 1]
            r = a[i + 1][j]
            s = a[i + 1][j + 1]

            # We need a unique identifier for each submatrix -
            matrices[f"{p}-{q}-{r}-{s}"] = 0

    return len(matrices)


for i in [i := input] * int(i()):
    l, _ = map(int, i().split())

    # Split input into a list
    m = [i().split() for j in range(l)]

    print(sub_matrices(m))
