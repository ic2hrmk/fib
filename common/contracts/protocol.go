package contracts

type ITransmitter interface {
	Send(Payload) error
}

type IReceiver interface {
	Receive() (Payload, error)
}

type Payload struct {
	NodeId    string
	FibNumber uint64
}