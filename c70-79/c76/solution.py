for i in range(int(input())):
    murderer_half_dna = input()

    people = []
    for i in range(int(input())):
        l = input()
        people.append(l)

    murderer_full_dna = []
    for p in murderer_half_dna:
        murderer_full_dna += [p, {"A": "T", "T": "A", "C": "G", "G": "C"}[p]]

    matches = []

    for member in people:
        dna, mem = member.split()
        matches += [[mem, len([1 for i in range(len(dna)) if dna[i] == murderer_full_dna[i]])]]

    highest = max([i[1] for i in matches])
    murderers = sorted([m[0] for m in matches if m[1] == highest], key=list)

    print(" ".join(murderers))
    
