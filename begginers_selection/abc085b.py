n = int(input())

list = []

for i in range(n):
  list.append(int(input()))

sorted_list = set(sorted(list, reverse=True))
print(len(sorted_list))
