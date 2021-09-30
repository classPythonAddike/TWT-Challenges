for i in[i:=input]*int(i()):
 l,_=map(int,i().split());a=[i().split()for j in range(l)];m={}
 for i in range(len(a)-1):
  for j in range(len(a[0])-1):m[f"{a[i][j]}-{a[i][j+1]}-{a[i+1][j]}-{a[i+1][j+1]}"] = 0
 print(len(m))
