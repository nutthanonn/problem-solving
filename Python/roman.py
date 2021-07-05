from collections import OrderedDict

def write_roman(num):

    roman = OrderedDict()
    roman[100] = "C"
    roman[90] = "XC"
    roman[50] = "L"
    roman[40] = "XL"
    roman[10] = "X"
    roman[9] = "IX"
    roman[5] = "V"
    roman[4] = "IV"
    roman[1] = "I"

    def roman_num(num):
        for r in roman.keys():
            x, y = divmod(num, r)
            yield roman[r] * x
            num -= (r * x)
            if num <= 0:
                break

    return "".join([a for a in roman_num(num)])


c,l,x,v,i = 0,0,0,0,0
for k in range(1, int(input()) + 1):
     a = write_roman(k)
     for j in a:
        if j == "C": c += 1
        elif j == "L": l += 1
        elif j == "X": x += 1
        elif j == "V": v += 1
        elif j == "I": i += 1
print(i, '', v, '', x, '', l, '', c)