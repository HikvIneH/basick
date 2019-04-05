def separ(i):
    digits = len(str(i))
    result = []
    for digit in range(digits):
        result.append(str(i)[digit] + '0' * (digits - digit -1))
    return '\n'.join(result)

x = 1345679
print(separ(x))