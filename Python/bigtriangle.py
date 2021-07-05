from itertools import combinations as CB
A = int(input()); a = [[int(a) for a in input().split()] for i in range(A)]
out = -1
b = [i for i in CB(a, 3)]
for j in b:
    x1 = j[0][0]; x2 = j[1][0]; x3 = j[2][0]
    y1 = j[0][1]; y2 = j[1][1]; y3 = j[2][1]
    c = abs((x1*y2) + (x2*y3) + (x3*y1) - (y1*x2) - (y2*x3) - (y3*x1))/2
    if out < c:
        out = c
print('%.3f' % out)
