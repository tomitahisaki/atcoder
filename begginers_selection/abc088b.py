n = int(input())
list = list(map(int, input().split()))
sorted_list = sorted(list, reverse=True)

alice = 0
bob = 0
for i in range(n):
  if i % 2 == 0:
    alice += sorted_list[i]
  else:
    bob += sorted_list[i]

print(alice - bob)
