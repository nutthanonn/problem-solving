price = input()
three = 0; eleven = 0
for n in price: three = (three * 10 + int(n)) % 3; eleven = (eleven * 10 + int(n)) % 11
print(three, eleven)
# idea from anonymous