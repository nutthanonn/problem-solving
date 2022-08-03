all_val = 0

for i in range(1_000, 9**5 * 9):
    val = sum([int(x)**5 for x in str(i)])
    if val == i:
        all_val += i
        print(val)
print(all_val)
