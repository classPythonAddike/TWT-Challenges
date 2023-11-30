import re
for i in[*open(0)][1:]:
 i=int(i)
 if i<9:print({1:1,4:2}.get(i,-1))
 else:print((p:=re.match(r"^0b10(?P<w>[1]*)0$",bin(i)))and 2+len(p.group(1))or-1)
