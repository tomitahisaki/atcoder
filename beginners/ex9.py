N = int(input())
A = int(input())

for i in range(N):
  op, B = input().split()
  if op == "+":
    A += int(B)
  if op == "-":
    A -= int(B)
  if op == "*":
    A *= int(B)
  if op == "/":
    if int(B) != 0:
      A //= int(B)
    else:
      print("error")
      break

  print(i + 1, A)
