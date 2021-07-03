import operator
from collections import Counter
a = int(input())
loop_like = [str(x) for x in input().split()]; numbeer = []
a = dict(Counter(loop_like))
A = [key for m in [max(a.values())] for key,val in a.items() if val == m]
for i in A: numbeer.append(int(i))
numbeer.sort()
for i in range(len(numbeer)):
    if i < len(numbeer) - 1: print(numbeer[i], end=" ")
    else: print(numbeer[i])