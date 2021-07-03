#test case
#  4 5
# A - - #
# # - B #
# A B # #
# # # # #
# 2
# 1 3 L
# 0 1 R
import timeit
start = timeit.default_timer()

n, m = 4,5
a = [["#A--#"], ["##-B#"], ["#AB##"], ["#####"]]
walk = [[1, 3, "L"], [0, 1, "R"]]

# R ขวา
# L ซ้าย
def move_box(position, lr, box1):
    box = [a for i in box1 for a in i]
    if lr == "R":
        if box[position + 1] == "-":
            box[position], box[position+1] = box[position+1], box[position]
    else:
        if box[position - 1] == "-":
            box[position], box[position-1] = box[position-1], box[position]
    return ["".join(box)]


def fall(x, m):
    c = []
    txt = [n for f in x for h in f for n in h]
    for i in range(len(txt)):
        if txt[i] == "A" or txt[i] == "B":
            if txt[i+m] == "-":
                txt[i], txt[i+m] = txt[i+m], txt[i]
    for i in range(0, len(txt), m):
        c.append(["".join(txt[i:i+m])])
    return c


def get_point(x, p, m):
    c = []
    point = p
    txt = [n for f in x for h in f for n in h]
    for i in range(len(txt)):
        if txt[i] == "A":
            if txt[i + m] == "A":
                txt[i+m], txt[i] = "-", "-"
                point += 5
            elif txt[i - 1] == "A":
                txt[i-1], txt[i] = "-", "-"
                point += 5
            elif txt[i + 1] == "A":
                txt[i+1], txt[i] = "-", "-"
                point += 5
        elif txt[i] == "B":
            if txt[i + m] == "B":
                txt[i+m], txt[i] = "-", "-"
                point += 5
            elif txt[i - 1] == "B":
                txt[i-1], txt[i] = "-", "-"
                point += 5
            elif txt[i + 1] == "B":
                txt[i+1], txt[i] = "-", "-"
                point += 5
    for i in range(0, len(txt), m):
        c.append(["".join(txt[i:i+m])])
    return c, point

    
    
# call function
point = 0
for i in walk:
    a[i[0]] = move_box(i[1], i[2], a[i[0]])
    for j in range(2):
        a = fall(a, m)
        a, point = get_point(a, point, m)
for i in a:
    for j in i:
        print(j)
print(point)


# check fall after replace '-' in function !



stop = timeit.default_timer()
print('Time: ', stop - start)  