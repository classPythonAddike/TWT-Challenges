for s in[*open(0)][1:]:
 m=0;s=s[:-1]
 for j in range(len(s)//2):m=[m,j+1][s[~j:]==s[:j+1]]
 print(m)
