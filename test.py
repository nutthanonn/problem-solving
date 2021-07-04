def move_box(position, lr, box1, point):
    box = [a for i in box1 for a in i]; check = "F"
    if lr == "R":
        if box[position] != "#" and box[position] != "-":
            if box[position + 1] == "-":
                box[position], box[position+1] = box[position+1], box[position]
                check = "T"
    elif lr == "L":
        if box[position] != "#" and box[position] != "-":
            if box[position - 1] == "-":
                box[position], box[position-1] = box[position-1], box[position]
                check = "T"
    if check == "F":
        point -= 5
    return ["".join(box)], point


def fall(x, m, c):
    z = []
    txt = [n for f in x for h in f for n in h]
    for j in c:
        for i in range(len(txt) - m - 1):
            if txt[i] == j:
                if txt[i+m] == "-":
                    txt[i], txt[i+m] = txt[i+m], txt[i]
    for i in range(0, len(txt), m):
        z.append(["".join(txt[i:i+m])])
    return z
def get_point(x, p, m, q, c):
    z = []
    point = p
    txt = [n for f in x for h in f for n in h]
    check = "T"
    for j in c:
        for i in range(len(txt) - m - 1):
            if txt[i] == j:
                if txt[i+1] == j and txt[i+2] == j and txt[i+m+1] == j:
                    txt[i], txt[i+2], txt[i+1], txt[i+m+1] = '-', '-', '-', '-'
                    point += 20
                    check = "F"
                elif txt[i+1] == j and txt[i+m] == j:
                    txt[i], txt[i+m], txt[i+1] = '-', '-', '-'
                    point += 15
                    check = "F"
                elif txt[i+1] == j and txt[i+2] == j:
                    txt[i+1], txt[i+2], txt[i] = '-', '-', '-'
                    point += 15
                    check = "F"
                elif txt[i+1] == j and txt[i+m+1] == j:
                    txt[i], txt[i+1], txt[i+m+1] = '-', '-', '-'
                    point += 15
                    check = "F"
                elif txt[i+1] == j:
                    txt[i], txt[i+1] = '-', '-'
                    point += 10
                    check = "F"
                elif txt[i+m] == j:
                    txt[i], txt[i+m] = '-', '-'
                    point += 10
                    check = "F"
    for i in range(0, len(txt), m):
        z.append("".join(txt[i:i+m]))
    return point, z

# input

# n, m = [int(a) for a in input().split()]
# a = []; walk = []; c = []
# for j in range(n):
#     b = input().split()
#     txt = ""
#     for i in range(len(b)):
#         txt += b[i]
#         if b[i] != '-' and b[i] != '#':
#             if b[i] not in c:
#                 c.append(b[i])
#     a.append([txt])
# for i in range(int(input())):
#     walk.append(input().split())
n, m = 5, 5
a = [['#B-A#'], ['##-##'], ['#A-A#'], ['#B-B#'], ['#####']]
walk = [[0, 1, "R"], [0, 1, "R"], [0, 3, "L"]]
c = ['A', 'B']
 # call function
point = 0
for i in walk:
    if int(i[0]) < n and int(i[1]) < m:
        a[int(i[0])], point = move_box(int(i[1]), str(i[2]), a[int(i[0])], point)
        for j in range(5):
            a = fall(a, m, c)
            point, a = get_point(a, point, m, j, c)




# print out!!!
print(point)
for i in a:
    for j in i:
        print(j, end=" ")
    print()
