class Tank:
    def __init__(self, name, ammo):
        self.name = name
        self.ammo = ammo
    
    def fire_ammo(self):
        if self.ammo > 0:
            self.ammo -= 1
    
    def add_ammo(self):
        if self.ammo < 20:
            self.ammo += 2

obj = Tank(input(), int(input()))
print(obj.ammo)
