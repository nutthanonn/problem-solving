import math
all_amicable_numbers = []


def find_divisors(num):
    final = []
    for i in range(2, math.floor(math.sqrt(num))):
        if num % i == 0:
            if num/i != i:
                final.append(i)
                final.append(num/i)

    final.append(1)
    return final


for a in range(10000):
    if a not in all_amicable_numbers:
        first = sum(find_divisors(a))
        secound = sum(find_divisors(first))
        if a == secound and a != first:
            all_amicable_numbers.append(a)
            all_amicable_numbers.append(first)

print(sum(all_amicable_numbers))
