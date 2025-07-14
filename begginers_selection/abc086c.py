n = int(input())
list = [list(map(int, input().split())) for _ in range(n)]
prev_t, prev_x, prev_y = 0, 0, 0
result = False

for t, x, y in list:
    dt = t - prev_t
    dist = abs(x - prev_x) + abs(y - prev_y)

    if dist > dt:
        result = False
        break
    elif (dt - dist) % 2 != 0:
        result = False
        break

    prev_t, prev_x, prev_y = t, x, y

    result = True

print("Yes" if result else "No")
