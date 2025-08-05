n = int(input())
s = str(input())

result = -1

for i in range(n - 2):
    if s[i : i + 3] == "ABC":
        result = i + 1
        break

print(result)
