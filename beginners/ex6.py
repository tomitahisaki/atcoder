A, op, B = input().split()
A = int(A)
B = int(B)

if op == "+":
  print(A + B)

# ここにプログラムを追記
elif op == "-":
  print(A - B)
elif op == "*":
  print(A * B)
elif op == "/":
  if B != 0:
    print(A // B)
  else:
    print("error")
else:
  print("error")
