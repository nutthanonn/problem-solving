from math import ceil as CE
N, M=  [int(a) for a in input().split()]
L,K = [int(b) for b in input().split()]
C = int(input())
miter = [[int(x) for x in input().split()] for i in range(N)]
price = 0
for i in range(N): price += sum(miter[i])
all_price = CE(((L*K*C) + price)/C); print(all_price)