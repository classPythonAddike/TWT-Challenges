# Challenge 73: Tekgar's late | Amogus #3

## Task
You are given a number T the number of tests cases. T cases follow. For each case you are given a number D for the total amount of doors, and on the next line a letter representing which way the lobby is (L for left, R for right, U for up, and D for down) and the next two lines arrays of digits representing Tekgar's and the cans' locations (0 for nothing at the door, 1 for a trashcan, and 2 for Tekgar). There will always be two arrays of digits of the exact same length. Figure out the minimum amount of steps (one step is movement from one door to one next to it) needed for tekgar to get to a trashcan which is preferably on his way to the lobby. To move from one hallway to the other (from one line of digits to the other) Tekgar must first get to the end of the current hallway (line of digits).

## Examples
### Input
```
2
6
R
012
000
8
D
1021
0010
```

### Output
```
1
1
```