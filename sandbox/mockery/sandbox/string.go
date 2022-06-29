package sandbox

type Stringer interface {
	String() string
}

type StringerImpl struct {
	id int
}

func (receiver StringerImpl) String() string {
	return "Hello from StringerImpl"
}