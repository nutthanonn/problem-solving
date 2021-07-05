from math import ceil as CE
ro = [[int(a) for a in input().split()] for i in range(int(input()))]
lis = [0]*5
for i in ro:
    lis[0] += i[0]*4
    lis[1] += i[1]*(3)
    lis[2] += i[2]*(2)
    lis[3] += i[3]*(1)
    lis[4] += i[4]*(1/2)
min_number = sum(lis)/4
print(CE(min_number))