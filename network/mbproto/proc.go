package mbproto

import (
	"io"
	"sync"
	"log"
	"github.com/golang/protobuf/proto"
	"crypto/cipher"
	"net"
	"encoding/binary"
)

const chanSize = 100

type Processor struct {
	needStop     int8
	stream       io.ReadWriter
	sendingChan  chan []byte
	receivedChan chan []byte
	stopChan     chan int
	waitGroup    sync.WaitGroup
	aead         *cipher.AEAD
	sessionKey   []byte
}

func NewProcessor(networkKey []byte, s io.ReadWriter) *Processor {
	p := &Processor{}
	p.sendingChan = make(chan []byte, chanSize)
	p.receivedChan = make(chan []byte, chanSize)
	return p
}

func (p *Processor) Handshake(isServer bool, networkKey []byte, privateIP net.IP) error {
	if !isServer {
		hm := HandshakeMessage{
			NetworkKey: networkKey,
			PrivateIP: uint32(privateIP.To4()),
		}
		_, err := p.writeMessage(hm)
		if err != nil {
			return err
		}
	}
	p.waitGroup.Add(2)
	go p.reader()
	go p.writer()
	return nil
}

func (p *Processor) writeMessage(msg proto.Message) (int, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		return 0, err
	}
	len := uint16(len(data))
	err = binary.Write(p.stream, binary.BigEndian, &len)
	if err != nil {
		return 0, err
	}
	n, err := p.stream.Write(data)
	return n + 2, err
}

func (p *Processor) readMessage(len int, msg proto.Message) (int, error) {
	var len uint16
	err := binary.Read(p.stream, binary.BigEndian, &len)
	if err != nil {
		return 0, err
	}
	io.ReadAtLeast()
	proto.Unmarshal()
}

func (p *Processor) reader() {
	defer p.waitGroup.Done()
	for {
		select {
		case <-p.stopChan:
			return
		}
	}
}

func (p *Processor) writer() {
	defer p.waitGroup.Done()
	for payload := range p.sendingChan {
		tm := &TransmitMessage{}
		tm.Payload = payload
		data, err := proto.Marshal(tm)
		if err != nil {
			log.Printf("proto marshal err: %s", err)
			return
		}

	}
}