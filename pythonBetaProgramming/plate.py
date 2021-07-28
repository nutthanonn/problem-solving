# NUT from KMUTT
C,S = [int(x) for x in input().split()]
sec1 = []; u = [[str(a) for a in input().split()] for i in range(S)]
for i in u: sec1.append([i[1], i[0]])
sec = dict(sec1)
q = []
while True:
    a = input().split()
    if a == list("X"): break
    else: q.append(a)
out = []

for i in range(len(q)):
    a = 0
    check = "T"; c1 = "T"; c2 = "T"
    if len(out) == 0 and q[i] == list("D"): print("empty")
    elif len(out) == 0: out.append(q[i])

    elif q[i] == list("D"): print(out[0][1]); out.pop(0); c2 = "F"
    else:
        for j in range(len(out)):
            if sec[q[i][1]] == sec[out[j][1]]:
                a = j + 1
                c1 = "F"
            elif c1 == "T":
                if sec[q[i][1]] != sec[out[j][1]]:
                    check = "F"
        if c1 == "T": out.append(q[i])
        else: out.insert(a, q[i])

if c2 == "F": print(0)