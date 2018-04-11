package server

import (
	"net"
	"log"
	"fmt"
	"bytes"
	"unsafe"
	"encoding/binary"

	protocol "github.com/ic2hrmk/fib/common/contracts"
)

type Server struct {
	hostAddress string
	rateLimiter *Limiter
	queue       *Queue
}

func NewServer(hostAddress string, flowSpeed float64, bufferSize int64) *Server {
	queue := NewQueue(bufferSize)

	rateLimiter := NewLimiter(int64(flowSpeed))
	go rateLimiter.Run()

	return &Server{
		hostAddress: hostAddress,
		rateLimiter: rateLimiter,
		queue:       queue,
	}
}

func (b Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", b.hostAddress)
	if err != nil {
		err = fmt.Errorf("failed to start listen at [%s], %s",
			b.hostAddress, err.Error(),
		)
		return err
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection %v", err)
			continue
		}

		if b.rateLimiter.IsAllowed() && b.queue.HasFreeSpace() {
			b.rateLimiter.AddRequest()
			go func() {
				err := b.handle(conn)
				if err != nil {
					log.Println("ERROR: ", err.Error())
				}
			}()
		}

	}
}

func (b Server) handle(conn net.Conn) (err error) {
	defer conn.Close()

	const payloadSize = int(unsafe.Sizeof(protocol.Payload{}))

	pbytes := make([]byte, payloadSize)
	_, err = conn.Read(pbytes[:])
	if err != nil {
		err = fmt.Errorf("failed to read transmission from [%s], %s",
			conn.RemoteAddr().String(), err.Error(),
		)
		return
	}

	payload := protocol.Payload{}

	binBuffer := bytes.NewBuffer(pbytes[:])
	err = binary.Read(binBuffer, binary.LittleEndian, &payload)
	if err != nil {
		err = fmt.Errorf("failed to read payload from from [%s], %s",
			conn.RemoteAddr().String(), err.Error(),
		)
		return
	}

	log.Println("PAYLOAD RECEIVED FROM [NODE_ID ", payload.NodeId, " ]")

	err = b.queue.Add(payload)
	if err != nil {
		log.Println("failed to add payload to queue,", err.Error())
	}

	return
}
