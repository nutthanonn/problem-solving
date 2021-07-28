from math import sqrt as SQ
number = int(input())
if number == 1 or number == 2: print('%.6f' % 2)
elif number == 3: print('%.6f' % (2 + SQ(3)))
elif (number % 2) == 0: print('%.6f' % number)
else: print('%.6f' % ((number - 3) + (SQ(3) * 2)))