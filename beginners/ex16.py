N, K = map(int, input().split())

list = list(map(str, input().split()))

result = [ item for item in list if len(item) >= K]
print(*result)
