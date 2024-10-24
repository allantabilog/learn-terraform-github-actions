package sandbox

type StringerClient struct {
	id       int
	stringer Stringer
}

func (receiver StringerClient) printMessage() {
	receiver.stringer.String()
}
