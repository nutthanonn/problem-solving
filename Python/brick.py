# Dek Chai Nut from Kmutt
N, M = [int(a) for a in input().split()]
state = ""
for i in range(N): state += input()
num = [int(c) for c in input().split()]

def run(x, m, n, t):
    po = x; r = list(t); s = 0
    for j in range(n):
        if j == 0 and r[po] == "#" or r[po] == "O": break
        elif r[x] == "O" or r[x] == "#": r[x - m] = "#"; break
        elif j == n-1 and r[x] == ".":r[x] = "#"; break
        x += m
    q = "".join(r)
    return q

for i in range(M):
    if num[i] != 0:
        for j in range(num[i]): state = run(i, M, N, state)
for i in range(1, M*N + 1):
    print(state[i - 1], end="")
    if (i%M) == 0: print()