import math


n = int(input())
i, j = 0, 0
s = int(math.sqrt(n))
UP, DOWN, LEFT, RIGHT = 1, 1, 1, 1

arr = [[0 for _ in range(s)] for _ in range(s)]


for k in range(1, s+1):
    arr[i][j] = k
    if i < s-1:
        i += 1

for k in range(s, s*2):
    arr[i][j] = k
    if j < s-1:
        j += 1

i -= 1

for k in range(s*2, s*3-1):
    arr[i][j] = k
    if i > 0:
        i -= 1

j -= 1
for k in range(s*3-1, s*4-3):
    arr[i][j] = k
    if j > 1:
        j -= 1

i, j = 1, 1

k = s*4-3
c = 0
while k != n + 1:
    if c == 0:
        while arr[i][j] == 0:
            arr[i][j] = k
            k += 1
            i += 1
        c += 1
        i -= 1
        j += 1
        DOWN += 1
    elif c == 1:
        while arr[i][j] == 0:
            arr[i][j] = k
            k += 1
            j += 1
        c += 1
        j -= 1
        i -= 1
        RIGHT += 1
    elif c == 2:
        while arr[i][j] == 0:
            arr[i][j] = k
            k += 1
            i -= 1
        c += 1
        i += 1
        j -= 1
        UP += 1
    else:
        while arr[i][j] == 0:
            arr[i][j] = k
            k += 1
            j -= 1
        c = 0
        j += 1
        i += 1
        LEFT += 1
for i in arr:
    for j in i:
        print(f"{str(j).zfill(3)}", end=" ")
    print()
print(f"UP : {UP}")
print(f"DOWN : {DOWN}")
print(f"LEFT : {LEFT}")
print(f"RIGHT : {RIGHT}")
