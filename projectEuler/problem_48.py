s = 0
mul = 10_000_000_000
for i in range(1, 1001):
    temp = i
    for j in range(1, i):
        temp *= i
        temp %= mul
    s += temp
    s %= mul
print(s)
