total = 0
name = []

with open("name.txt", encoding='utf-8') as f:
    name = f.read().split("\",\"")

l = len(name) - 1
name[0] = name[0][1:]
name[l] = name[l][:-1]
name.sort()
for i in range(len(name)):
    temp = 0
    for j in name[i]:
        temp += ord(j) - 64
    total += (i+1)*temp

print(total)
