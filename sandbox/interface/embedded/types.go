package embedded

type People interface {
	GetName() string
	GetAge() int
}

type Student interface {
	People
	GetScore() float64
	GetSchoolName() string
}

type StudentImpl struct {
	name   string
	age    int
	score  float64
	school string
}

func NewStudent() Student {
	var s = new(StudentImpl)
	s.name = "Pepe"
	s.age = 8
	s.score = 85.00
	s.school = "Coburg High"
	return s
}

func (a *StudentImpl) GetName() string {
	return a.name
}

func (a *StudentImpl) GetAge() int {
	return a.age
}

func (a *StudentImpl) GetScore() float64 {
	return a.score
}

func (a *StudentImpl) GetSchoolName() string {
	return a.school
}
