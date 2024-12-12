package impl1

type IncrementerImpl1 struct {
	value int
}

func NewIncrementerImpl1() *IncrementerImpl1 {
	return &IncrementerImpl1{}
}

func (i *IncrementerImpl1) Increment() int {
	i.value++
	return i.value
}