B=[0];a=[1]*951161;a[:1]=0,0
for(i,p)in enumerate(a):
 if p:B+=[i];a[i*i::i]=[0]*len(a[i*i::i])
for i in[*open(0,'r')][1:]:print(B[int(i)])
