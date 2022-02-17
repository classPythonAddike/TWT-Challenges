for i in[I:=input]*int(I()):
	n=[*map(int,I().split())];l=len(n);d={i:0for i in range(l)};p=0
	while all(i<2for i in d.values()):
		p=(p+n[p])%l;d[p]+=1
		if all(d.values()):print(n!=[0]);break
	else:print(False)