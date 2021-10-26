# Challenge 76: Who's the murderer? | Amogus #6

## Task
Given T number of testcases, you need to accept -
1. A half-DNA strand found at the crime scene (say this is murder_dna)
2. n, an integer denoting how many members are onboard
3. For n lines - the full DNA strand of every staff member on board, and their names (say this is dna_list):cheemsadidas~2:

You need to construct the full DNA strand of the murderer (will explain more about this), and then compare the full DNA strand of murder_dna with the full DNA strand of every staff member. The member who shares the most full DNA with the murderer is the prime suspect. Print the name of the prime suspect.

DNA is a double helix (essentially, two ropes held together with bonds) is made up of four subunits - Adenine, Thymine, Guanine, Cytosine. A and T always bond together, and C and G always bond together, thus holding DNA together. You will be given the subunits on one of the "ropes", and you need to construct the other half, by inserting each character's bonding pair into the DNA given (A -> T will be inserted, G -> C will be inserted, C -> G will be inserted, and T -> A will be inserted).

#### Example
if you are given ACGAT as a half strand, then -
1. A -> Insert a T next to it
2. C -> insert a G next to it
3. G -> Insert an C next to it
4. A -> Insert a T next to it
5. T -> Insert a A next to it
And you get ATCGGCATTA 

### Sample Input:
```
1                         -> There is 1 test case
ACGA                      -> This is the half strand of DNA found at the murder scene
3                         -> 3 people onboard
AGGTAAAT Espiobest        -> Full strand of Espiobest's DNA
AGAGAGTC Tekgar           -> Full strand of tekgar's DNA
CCCTAAAG Avib             -> Full strand of Avib's DNA
```

### Output
```
Espiobest
```

### Explanation
```
Murderer's full DNA strand   - ATCGGCAT
Espiobest's full DNA strand  - AGGTAAAT
Tekgar's full DNA strand     - AGAGAGTC
Avib' full DNA strand        - CCCTAAAG
```

Espiobest shares 3 characters in the same position with the murderer (A, A, T)
Tekgar shares 2 characters in the same position with the murderer (A, G)
Avib shares 2 characters in the same position with the murderer (C, A)

So Espiobest is the murderer

In case two members share the same number of characters with the murderer's DNA, return both of them, arranged alphabetically (Why can't we have more than one murderer)

The length of the full DNA strand of each member will always be the same, and will be equal to 2 * length of murderer's half DNA strand

