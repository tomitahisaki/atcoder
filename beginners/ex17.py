N, M = map(int, input().split())
AB = [list(map(int, input().split())) for i in range(M)]

# ここにコードを追記する
result = [["-"] * N for _ in range(N)]

for a, b in AB:
  result[a - 1][b - 1] = "o"
  result[b - 1][a - 1] = "x"

print(*[" ".join(row) for row in result], sep="\n")
