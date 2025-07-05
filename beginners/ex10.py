N = int(input())
# ここにプログラムを追記

scores = list(map(int, input().split()))
result = 0

for i in range(N):
  result += scores[i - 1]

average = result // N

for i in range(N):
  if scores[i] >= average:
    print(scores[i] - average)
  else:
    print(average - scores[i])
