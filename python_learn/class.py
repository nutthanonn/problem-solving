#สร้าง Class ง่ายๆ
class school:
    def __init__(self, name, userID):
        self._name = name
        self.__userID = userID

    def _showData(self):
        print("Name : {}".format(self._name))
        print("USER-ID : {}".format(self.__userID))

class Teacher(school):
    def __init__(self, name, userID, subject):
        super().__init__(name, userID)
        self.subject = subject
        
    def _showData(self):
        super()._showData()
        print("Subject : {}".format(self.subject))


class student(school):
    def __init__(self, name, userID, section):
        super().__init__(name, userID)
        self.section = section

    def _showData(self):
        super()._showData()
        print("Section : {}".format(self.section))


obj1 = Teacher("nut", 1150, "math")
obj1._showData()
