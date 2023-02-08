package student

import "fmt"

type Student struct {
	Name       string
	Age, Grade int
}

func NewStudent(name string, age, grade int) *Student {
	return &Student{Name: name, Age: age, Grade: grade}
}

func (s *Student) GetName() string {
	return s.Name
}

func (s *Student) String() string {
	return fmt.Sprintf("Student | Name: %10s, age: %2d, grade: %2d", s.Name, s.Age, s.Grade)
}
