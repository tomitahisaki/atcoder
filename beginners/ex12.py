N = int(input())  # 生徒の数Nを読み込む
T = list(map(int, input().split()))  # 各生徒のゴールまでの時間を読み込む
# ここにプログラムを追記

fastest_time = min(T)

for i, v in enumerate(T):
  if v == fastest_time:
    print(i + 1)  # 生徒の番号は1から始まるので、i + 1を出力
