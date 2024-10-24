package sandbox

import "fmt"

type Iff interface {
	M()
}

type Tff struct {
	S string
}

func (t *Tff) M() {
	fmt.Println(t.S)
}

func InterfaceMain() {
	var i Iff = &Tff{"hello"}
	i.M()
}
