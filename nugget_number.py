a = int(input())
number = []

if a < 6:
    print("no")
else:
    for i in range(0, 17):
        for j in range(0, 13):
            for k in range(0, 6):
                if 6*i + 9*j + 20*k not in number:
                    number.append(6*i + 9*j + 20*k)
    number.sort()
    number.pop(0)
    for i in number:
        if i <= a: print(i)
        else: break