# sol = input()
# number = [[int(a) for a in input().split()[:-1:]] for i in range(int(input()))]



# test case
# a*b^c+d*e^f
# 2
# 1 0
# 2 0
sol = 'a*b^c+d*e^f'
number = [[1, 0], [2, 0]]

sol1 = '(a+b)+e*h+g*w^e+(f^gh)'


def cut_solution_math(x):
    b = list(x)
    n, m = 0, 0
    txt = ''; txt1 = ''
    ls = []
    ch = 0; ch1 = 0
    for i in range(len(b)):
        if b[i] == '(': n = i
        elif b[i] == ')': m = i
        
        if m != 0:
            for i in range(n, m + 1):
                txt += b[i]
            ls.append(txt)
            txt = ''; n, m = 0, 0

        elif n == 0 or m == 0:
            if ch <= 2:
                if b[i] == "+":
                    ch1 += 1
                    n1 = i
                    ch += 1
            else:
                n2 = n1 - ch1
                for i in range(n2, n1):
                    txt1 += b[i]
                ls.append(txt1)
                n1, ch, ch1 = 0, 0, 0
                txt1 = ""
    

                
                











    return ls



a = cut_solution_math(sol1)
print(a)