for _ in[y:=input]*int(y()):u="AGCT";f=sum([[x,u[~u.index(x)]]for x in y()],[]);t={e:sum(o==f[i]for i,o in enumerate(d))for d,e in map(str.split,map(input,int(y())*[""]))};print(*sorted(m for m in t if t[m]==max(t.values())))