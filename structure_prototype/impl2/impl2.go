package impl2

type IncrementerImpl2 struct {
	value int
}

func NewIncrementerImpl2() *IncrementerImpl2 {
	return &IncrementerImpl2{}
}

func (i *IncrementerImpl2) Increment() int {
	i.value+=10
	return i.value
}

