from itertools import permutations as pmu
type = int(input()); q = int(input())
dont = [int(a) for a in input().split()]; food_ = []; out = []
for i in range(1, type + 1): food_.append(i)
food  = [x for x in pmu(food_)]
for i in food:
    if i[0] not in dont:
        out.append(i)
for i in out:
    for j in range(len(i)):
        if j < len(i) - 1: print(i[j], end=" ")
        else: print(i[j])