for i in[i:=input]*int(i()):
	r,d,h,n=range,len,print,i()
	if len(n)==1:h(False);break
	for l in r(1,d(n)//2+1):
		s,q=int(n[:l]),0
		if d(n)%l:continue
		for w in r(s,s+d(n)//l):	
			if str(w)!=n[(w-s)*l:(w-s)*l+l]:q=0;break
			q=1
		if q:h(True);break
	else:h(False)