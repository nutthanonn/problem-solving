score = {'a':1, 'e':1, 'i':1, 'o':1, 'u':1, 'n':1, 'r':1, 't':1, 'l':1, 's':1, 'd':2, 'g':2, 'b':3, 'c':3, 'm':3, 'p':3, 'f':4, 'h':4, 'v':4, 'w':4, 'y':4, 'k':5, 'j':8, 'x':8, 'q':10, 'z':20}

out = 0


def check(txt, score):
    num = 0
    for i in txt:
        for j in score:
            if i == j:
                num += score[j]
    return num


txt = ['some', 'first', 'potsie', 'day', 'cloud', 'have', 'back', 'this']

# for i in range(int(input())):
#     txt = input()
for i in txt:
    if out < check(i, score):
        out = check(i, score)
        word = i

print(word)