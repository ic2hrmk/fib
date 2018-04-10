package transmitter

import (
	"net"
	"fmt"
	"time"
	"bytes"
	"encoding/binary"

	"github.com/ic2hrmk/fib/common/contracts"
	"github.com/ic2hrmk/fib/generator/config"
)

type Transmitter struct {
	conn net.Conn
}

func NewTransmitter() (transmitter *Transmitter, err error) {
	transmitter = &Transmitter{}
	loggerAddress := config.GetLoggerAddress()

	transmitter.conn, err = net.Dial("tcp", loggerAddress)
	if err != nil {
		switch err.(type) {
		case *net.OpError:
			err = fmt.Errorf(
				"failed to connect to logger at [%s], %s",
				loggerAddress, err.Error(),
			)
		default:
			err = fmt.Errorf(
				"unknown error during connection establishing to logger at [%s]: %s",
				loggerAddress, err.Error(),
			)
		}

		return
	}

	return
}

func (t Transmitter) Send(payload contracts.Payload) (err error) {
	//	Limit write deadline
	t.conn.SetWriteDeadline(time.Now().Add(1 * time.Second))

	//	Data serialization
	var binBuffer bytes.Buffer
	binary.Write(&binBuffer, binary.LittleEndian, payload)

	_, err = t.conn.Write(binBuffer.Bytes())
	if err != nil {
		return
	}

	return
}