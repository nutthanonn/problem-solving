a = []
while True:
    x = int(input())
    if x == 0: break
    else: a.append(x)
text = input().lower()
if text == "min": a.sort()
elif text == "max": a.sort(reverse = True)
for i in a: print(i, end= " ")