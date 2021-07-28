a, b ,c = [int(x) for x in input().split()]
po = []; fif = 0; zero = 0
for i in range(c): po.append(input().split())
position = [0]*(a)
for i in po:
    if int(i[0]) + int(i[1]) > a: l = a
    else: l = int(i[0]) + int(i[1])
    for j in range(int(i[0]), l):
        position[j] += 1
for i in position:
    if i == 0: zero += b
    elif i == 1: fif += b
print(str(zero) + " " + str(fif))