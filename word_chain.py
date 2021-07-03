len_word = int(input()); ro = int(input()); word = []
for i in range(ro): word.append(input())

def check_word(a, b):
    count = 0
    for i in range(len_word):
        if a[i] != b[i]: count += 1
    return True if count >= 3 else False

for i in range(ro):
    if i < ro - 2:
        if check_word(word[i], word[i+1]):
            print(word[i])
            break
    elif i == ro-1: print(word[i])