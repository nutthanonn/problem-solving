name = "1729. Find Followers Count"

name = name.split()
if len(name[0]) != 5:
    name[0] = name[0].zfill(5)
name[0] = name[0][:-1] + "_"
name = name[0] + "-".join(name[1:])
print(name)
