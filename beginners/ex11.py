S = input() # ここにプログラムを追記

resuilt = 1

for i in range(len(S)):
  if S[i] == "+":
    resuilt += 1
  elif S[i] == "-":
    resuilt -= 1
  else:
    continue

print(resuilt)
