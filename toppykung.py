ro = int(input()); name = []; name1 = []
for i in range(ro): name.append(input())
for i in name: 
    if i not in name1: name1.append(i)
name1.sort()
for i in name1: print(i)