A,B,C = [int(x) for x in input().split()]; out = []
for a in range(1, A+1):
    for b in range(-abs(C), abs(C) + 1):
        if b != 0 and a*(A//a) == A and b*(C//b) == C and a*(C//b) + b*(A//a) == B:
            out.append([a, b, A//a, C//b])
if len(out) < 1: print("No Solution")
else: a = min(out); print(a[0], '', a[1], '', a[2], '', a[3])