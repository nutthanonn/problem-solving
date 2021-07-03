number = [int(a) for a in input().split()]; prime = []; num_sort = []
for i in range(2, 1000):
    if i > 1:
        for j in range(2, i):
            if (i%j) == 0:
                break
        else:
            prime.append(i)
for i in prime:
    if i > number[0]:
        break
    for j in range(2, number[0] + 1):
        if (j%i) == 0:
            if j not in num_sort:
                num_sort.append(j)
print(num_sort[number[1] - 1])
