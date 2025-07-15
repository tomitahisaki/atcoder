# Imos method 差分配列
n, m = map(int, input().split())

list = [0 for _ in range(n + 1)]

for i in range(m):
    left, right = map(int, input().split())
    list[left - 1] += 1  # add 1 at start index
    list[right] -= 1  # add -1 at end index

for i in range(1, n + 1):
    list[i] += list[i - 1]  # culculate cumulative sum

list = list[:n]
print(min(list))
