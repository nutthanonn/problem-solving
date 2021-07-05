ro = int(input()); number = [int(a) for a in input().split()]
# ro = 8; number =[4, -6, 3, -2, 6, -4, -6, 6]
check = 0; out = 0

# find max_number
for i in range(ro):
    check = 0
    for j in range(i, ro):
        check += number[j]
        if check > out: out = check; po, po1 = i, j + 1

# cut list
if out != 0:
    for i in range(po, po1):
        if i < po1 - 1: print(number[i], end=" ")
        else: print(number[i])
    print(out)
else: print("Empty sequence")