m = int(input())
n = int(input())
number = []
for i in range(m*n): number.append(float(input()))
under = max(number) - min(number)
position = 0
for i in range(m):
    for j in range(n):
        print("%.4f" % ((number[position]-min(number))/under), end=" ")
        position += 1
    print()