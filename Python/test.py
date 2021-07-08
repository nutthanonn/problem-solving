from math import comb
N, M = [int(a) for a in input().split()]
out = 0
check = "T"
def intiger(n, r):
    return comb(n, r)


if  N<M:
    print(1)
    check = "F"
for i in range(N//M + 1):
    n = N-(i*M) + i
    r = i
    out += intiger(n, r)
if check == "T":
    a = divmod(out, 1000000007)
    print(a[1])
