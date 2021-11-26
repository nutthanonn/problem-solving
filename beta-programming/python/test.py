b = []
a = ""
while a != ['end']:
    a = input().split(",")
    check = True
    if len(b) == 0:
        b.append(a)
        check = False
    else:
        for i in b:
            if i[0] == a[0]:
                i[1] = int(i[1]) + int(a[1])
                check = False
    if check:
        b.append(a)
    print(b)
