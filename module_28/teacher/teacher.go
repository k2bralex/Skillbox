package teacher

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	Salary float32
}

func NewTeacher(name string, age int, salary float32) *Teacher {
	return &Teacher{Name: name, Age: age, Salary: salary}
}

func (t *Teacher) GetName() string {
	return t.Name
}

func (t *Teacher) String() string {
	return fmt.Sprintf("Teacher | Name: %10s, age: %2d, salary: %10.2f", t.Name, t.Age, t.Salary)
}
