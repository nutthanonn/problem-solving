fac = [0]*1001
fac[0] = 1
for i in range(1, 101):
    fac[i] = fac[i-1]*i

print(sum([int(x) for x in list(str(fac[100]))]))
