import math


triangle_numbers = [0]*100000
triangle_numbers[0] = 1
divisors = []

i = 1
check = True
while check:
    triangle_numbers[i] = triangle_numbers[i-1] + (i+1)
    for j in range(1, math.floor(math.sqrt(triangle_numbers[i])) + 1):
        if triangle_numbers[i] % j == 0:
            if triangle_numbers[i]/j != j:
                divisors.append(j)
                divisors.append(triangle_numbers[i]/j)

    if(len(divisors) >= 500):
        check = False
    else:
        divisors.clear()
    i += 1

print(triangle_numbers[i-1])
