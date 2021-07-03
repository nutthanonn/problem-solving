round = int(input()); a = []
for j in range(round): a.append(int(input()))
for i in a:
    if i == 2: print("T")
    elif (i%2) == 0: print("F")
    else: print("T")
