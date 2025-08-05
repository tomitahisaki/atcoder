h, w, d = map(int, input().split())

list = [list(input()) for _ in range(h)]

floor_list = []

for i in range(h):
    for j in range(w):
        if list[i][j] == ".":
            floor_list.append((i, j))  # list of  floor positions

N = len(floor_list)
answer = 0

for ii in range(N):
    for jj in range(ii + 1, N):  # brute force search all pairs of floors
        i1, j1 = floor_list[ii]
        i2, j2 = floor_list[jj]
        cnt = 0
        for i in range(h):
            for j in range(w):
                if (
                    list[i][j] == "."
                    and min(abs(i - i1) + abs(j - j1), abs(i - i2) + abs(j - j2)) <= d
                ):
                    cnt += 1
        answer = max(answer, cnt)

print(answer)
