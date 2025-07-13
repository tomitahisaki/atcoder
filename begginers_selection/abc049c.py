s = str(input())

while True:
    if s.endswith("dream"):
        s = s[:-5]
    elif s.endswith("dreamer"):
        s = s[:-7]
    elif s.endswith("erase"):
        s = s[:-5]
    elif s.endswith("eraser"):
        s = s[:-6]
    else:
        break

if s == "":
    print("YES")
else:
    print("NO")
