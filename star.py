A = int(input())
def even(a):
    num = a//2 - 1; rou = a//2; txt = 0; even_list = []
    for i in range(rou):
        txt = i*2 - 1
        if i == 0: even_list.append("-"*num +  "*"  +  "-"*num)
        else: even_list.append("-"*num +  "*" +  "-"*txt + "*" + "-"*num)
        num -= 1
    for i in even_list: print(i)
    for j in reversed(even_list): print(j)

def odd(b):
    odd_list = []; round = int(b/2) + 1; text = 0; number = (b-1)//2
    for i in range(round):
        text = i*2 - 1
        if i == 0: odd_list.append("-"*number + "*" + "-"*number)
        else: odd_list.append("-"*number + "*" + "-"*text + "*" + "-"*number)
        number -= 1
    for i in odd_list: print(i)
    odd_list.pop(len(odd_list) - 1)
    for j in reversed(odd_list): print(j)

if (A%2) == 0: even(A)
else: odd(A)