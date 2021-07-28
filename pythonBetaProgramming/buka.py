a = []
for i in range(3): a.append(input().strip())
if a[1] == "+": print(int(a[0]) + int(a[2]))
elif a[1] == "*": print(int(a[0]) * int(a[2]))