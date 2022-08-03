from num2words import num2words


def num_to_word():
    result = 0
    for i in range(1, 1001):
        result += (len(''.join(''.join(num2words(i).split('-')).split(' '))))
    return result


print(num_to_word())
