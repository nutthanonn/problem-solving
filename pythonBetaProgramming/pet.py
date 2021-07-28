score = [[int(a) for a in input().split()] for i in range(5)]; score_all = []
rank = 0
for i in score:
    s = 0
    for j in i:
        s += j
    score_all.append(s)
sc = max(score_all)
for i in score_all:
    rank += 1
    if i == sc:
        break
print(str(rank) + " " + str(sc))