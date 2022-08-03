long_chain = 1
curr = 0
r = 0

for i in range(1, 1000000):
    r += 1
    chain = 1
    while i != 1:
        if(i % 2 == 0):
            i /= 2
        else:
            i = 3*i+1
        chain += 1

    if chain >= long_chain:
        long_chain = chain
        curr = r


print(curr)
