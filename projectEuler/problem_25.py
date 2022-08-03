fib = {
    1: 1,
    2: 1
}

i = 3
n = 0
while len(str(n)) != 1000:
    n = fib[i-1] + fib[i-2]
    fib[i] = n
    i += 1
print(i-1)
