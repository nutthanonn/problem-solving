def sum_prime(num):
    summ = 0
    for i in range(3, num, 2):
        for e in range(2, int(i ** 0.5) + 1):
            if not i % e:
                break
        else:
            summ += i
    return summ


print(sum_prime(2_000_000))
