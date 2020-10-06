res = []

current = '012'

n = 10
m = 3

while current != '789':
    print(current)
    while int(current[2]) < (n - m + 2):
        current = current[:2] + str(int(current[2]) + 1)
        print(current)

    while int(current[1]) < (n - m + 1):
        current = current[0] + str(int(current[1]) + 1) + str(int(current[1]) + 2)
        print(current)

        while int(current[2]) < (n - m + 2):
            current = current[:2] + str(int(current[2]) + 1)
            print(current)

    while int(current[0]) < (n - m):
        current = str(int(current[0]) + 1) + str(int(current[0]) + 2) + str(int(current[0]) + 3)
        print(current)


        while int(current[1]) < (n - m + 1):
            current = current[0] + str(int(current[1]) + 1) + str(int(current[1]) + 2)
            print(current)

            while int(current[2]) < (n - m + 2):
                current = current[:2] + str(int(current[2]) + 1)
                print(current)
