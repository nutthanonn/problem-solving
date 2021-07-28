from datetime import datetime as dt
d = input().split()
Day = dt(2009, int(d[1]), int(d[0]))
print(Day.strftime(("%A")))