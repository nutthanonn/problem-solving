from math import ceil as CE
num = [int(a) for a in input().split()]
if num[0] > num[1]: print(2)
else: print(CE(num[1] / num[0]))