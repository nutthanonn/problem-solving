Ad = "ABC"*100
Bru = "BABC"*100
Gor = "CCAABB"*100
a = 0; b = 0; g = 0
number = int(input())
ans = input()
for i in range(number):
    if ans[i] == Ad[i]:
        a += 1
    if ans[i] == Bru[i]:
        b += 1
    if ans[i] == Gor[i]:
        g += 1
score = max(a,b,g)
print(score)
if score == a:
    print("Adrian")
if score == b:
    print("Bruno")
if score == g:
    print("Goran")