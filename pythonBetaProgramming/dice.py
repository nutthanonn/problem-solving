# Nut from KMUTT
ro = int(input())
spin = [[str(a) for a in input()]for i in range(ro)]
c = []

def dice(x, a):
    if x == "F": txt = a[3] + a[0] + a[2] + a[5] + a[4] + a[1]
    elif x == "B": txt = a[1] + a[5] + a[2] + a[0] + a[4] + a[3]
    elif x == "L": txt = a[4] + a[1] + a[0] + a[3] + a[5] + a[2]
    elif x == "R": txt = a[2] + a[1] + a[5] + a[3] + a[0] + a[4]
    elif x == "C": txt = a[0] + a[4] + a[1] + a[2] + a[3] + a[5]
    elif x == "D": txt = a[0] + a[2] + a[3] + a[4] + a[1] + a[5]
    return txt


for i in spin:
    b = "123546"
    for j in i:
        for k in j:
            b = dice(k, b)
    c.append(b[1])
print(" ".join(c))
