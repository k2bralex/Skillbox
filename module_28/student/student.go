package student

type Student struct {
	Name       string
	Age, Grade int
}

func NewStudent(name string, age, grade int) *Student {
	return &Student{Name: name, Age: age, Grade: grade}
}
