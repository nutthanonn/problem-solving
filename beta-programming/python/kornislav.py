number = [int(i) for i in input().split()]
number.sort(reverse=True)
print(number[1]*number[3])