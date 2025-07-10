n, a, b = map(int, input().split())

ans = 0

for i in range(1, n + 1):
  sum = 0
  for j in str(i):
    sum += int(j)
  if a <= sum <= b:
    ans += i

print(ans)
