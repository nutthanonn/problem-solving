text = input()
dic = {"apa":"a", "epe":"e", "ipi":"i", "opo":"o", "upu":"u"}
for i in dic:
    text = text.replace(i,dic[i])
print(text)
