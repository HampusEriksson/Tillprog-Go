class Student:
    def __init__(self, name, age, grade):
        self.name = name
        self.age = age
        self.grade = grade

    def display_student_info(self):
        print(f"Name: {self.name}, Age: {self.age}, Grade: {self.grade}")

# Creating instances of the Student class
student1 = Student("Alice", 17, "A")
student2 = Student("Bob", 18, "B")

# Displaying student information
student1.display_student_info()
student2.display_student_info()