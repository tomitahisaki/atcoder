n = input()
list = [int(i) for i in input().split()]

count = 0

def devide_list(list):
  global count
  odd_number_exist = False

  for i in range(len(list)):
    if list[i] % 2 == 0:
      list[i] = list[i] // 2
    else:
      odd_number_exist = True
      break

  if odd_number_exist:
    return count
  else:
    count += 1
    devide_list(list)

devide_list(list)
print(count)

