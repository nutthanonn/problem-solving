# Nut from KMUTT
def get_number(x):
    if x == "     |  |": n = 1
    elif x == " _  _||_ ": n = 2
    elif x == " _  _| _|": n = 3
    elif x == "   |_|  |": n = 4
    elif x == " _ |_  _|": n = 5
    elif x == " _ |_ |_|": n = 6
    elif x == " _   |  |": n = 7
    elif x == " _ |_||_|": n = 8
    elif x == " _ |_| _|": n = 9
    elif x == " _ | ||_|": n = 0
    return str(n)



n,m = [int(a) for a in input().split()]
a = []; b = []
num = ""; num1 = ""
for i in range(3):
    a1 = input() + " "
    for j in range(len(a1)):
        if (j%4) == 3:
            a.append(a1[j-3:j])
for i in range(3):
    b1 = input() + " "
    for j in range(len(b1)):
        if (j%4) == 3:
            b.append(b1[j-3:j])

for i in range(len(a) - n*2):
    txt = ""
    for x in a[i]: txt += x
    for x in a[i + n]: txt += x
    for x in a[i + n*2]: txt += x
    num += get_number(txt)

for i in range(len(b) - m*2):
    txt1 = ""
    for y in b[i]: txt1 += y
    for y in b[i + m]: txt1 += y
    for y in b[i + m*2]: txt1 += y
    num1 += get_number(txt1)

print(int(num) + int(num1))