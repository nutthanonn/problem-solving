a = [int(a) for a in input().split()]; number = []
for i in a:
    if (i%2) == 0 or i == 1: number.append(i)
    else: number.append(i - 1)
count = 0
for i in number:
    num = i
    while num != 1:
        if (num%2) == 1:
            num -= 1
        num /= 2
        count += 1
print(count)