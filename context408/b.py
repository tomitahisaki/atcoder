n = int(input())
list = map(int, input().split())

unique_sorted_list = sorted(set(list))

print(len(unique_sorted_list))
print(*unique_sorted_list)
