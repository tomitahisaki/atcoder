n, s = map(int, input().split())
list = map(int, input().split())

prev_time = 0
sleep_time = s + 0.5
result = False
for i in list:
    if sleep_time > i - prev_time:
        prev_time = i
    else:
        result = False
        break

    result = True

print("Yes" if result else "No")
