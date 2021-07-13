a = [int(a) for a in input().split()]
for i in a:
    if i == -1: break
    else:
        work_bee = 1
        soilder_bee = 0
        host = 1
        for i in range(i):
            new_work_bee = work_bee + soilder_bee + host
            soilder_bee = work_bee
            work_bee = new_work_bee
            a = work_bee + soilder_bee + host
        print(work_bee, a)