n, m = map(int, input().split())

list_a = list(map(int, input().split()))
list_b = list(map(int, input().split()))

result = 0
for i in list_b:
    result += list_a[i - 1]

print(result)
