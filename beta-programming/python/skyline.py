r = int(input())
n = [[int(a) for a in input().split()] for i in range(r)]
ro = [0]*300
for i in n:
    a = i[0]; b = i[1]; c = i[2]
    for j in range(a, c):
        if ro[j] < b: ro[j] = b
for i in range(300):
    if ro[i] != ro[i - 1]:
        print(i, ro[i], end=" ")