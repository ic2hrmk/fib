package transmitter

import (
	"net"
	"fmt"
	"sync"
	"time"
	"bytes"
	"encoding/binary"

	"github.com/ic2hrmk/fib/common/contracts"
)

type Transmitter struct {
	mutex sync.Mutex

	identification  [8]byte
	receiverAddress string
}

func NewTransmitter(identification [8]byte, receiverAddress string) *Transmitter {
	return &Transmitter{
		identification: identification,
		receiverAddress: receiverAddress,
	}
}

func (t Transmitter) Send(payload contracts.Payload) (err error) {
	//	Make transmitter thread-safe
	t.mutex.Lock()
	defer t.mutex.Unlock()

	//	Identify myself as sender
	copy(payload.NodeId[:], t.identification[:])

	//	Data serialization
	var binBuffer bytes.Buffer
	err = binary.Write(&binBuffer, binary.LittleEndian, &payload)
	if err != nil {
		err = fmt.Errorf(
			"failed to serialize payload, %s", err.Error(),
		)
		return
	}

	var conn net.Conn
	conn, err = net.Dial("tcp", t.receiverAddress)
	if err != nil {
		switch err.(type) {
		case *net.OpError:
			err = fmt.Errorf(
				"failed to connect to receiver at [%s], %s",
				t.receiverAddress, err.Error(),
			)
		default:
			err = fmt.Errorf(
				"unknown error during connection establishing to receiver at [%s]: %s",
				t.receiverAddress, err.Error(),
			)
		}

		return
	}

	//	Limit write deadline
	conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
	_, err = conn.Write(binBuffer.Bytes())
	if err != nil {
		err = fmt.Errorf(
			"failed to transmit data to receiver at [%s]: %s",
			t.receiverAddress, err.Error(),
		)
		return
	}

	return
}