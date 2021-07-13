class school:
    def __init__(self, teacher, student, name):
        self.teacher = teacher
        self.student = student
        self.name = name

    def show(self):
        print("Name : " + self.name)




class Teacher(school):
    def __init__(self, department, salary, name):
        self.department = department
        self.salary = salary
        self.name = name
        super().__init__(self.name)
        super().show()


obj1 = Teacher("math", 40000, "NUTTHANON")

