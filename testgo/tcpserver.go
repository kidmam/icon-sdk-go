package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
)

func init() {
	gob.Register(MyMsgBodyPing{})
}

func main() {
	l, err := net.Listen("tcp", ":5032")
	if nil != err {
		log.Fatalf("fail to bind address to 5032; err: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Printf("fail to accept; err: %v", err)
			continue
		}
		go ConnHandler(conn)
	}
}

type MyMsg struct {
	Header MyMsgHeader
	Body   interface{}
}
type MyMsgHeader struct {
	MsgType string
	Date    string
}
type MyMsgBodyPing struct {
	Content string
}

type MyConnection struct {
	conn        net.Conn
	dec         *gob.Decoder
	codecBuffer bytes.Buffer
	recvBuffer  []byte
}

func (mc *MyConnection) RecvMessage() (MyMsg, error) {
	lengthBuf := make([]byte, 4)
	_, err := mc.conn.Read(lengthBuf)
	if nil != err {
		return MyMsg{}, err
	}
	msgLength := binary.LittleEndian.Uint32(lengthBuf)

	mc.codecBuffer.Reset()

	for 0 < msgLength {
		n, err := mc.conn.Read(mc.recvBuffer)
		if nil != err {
			return MyMsg{}, err
		}
		if 0 < n {
			data := mc.recvBuffer[:n]
			mc.codecBuffer.Write(data)
			msgLength -= uint32(n)
		}
	}

	msg := MyMsg{}
	if err = mc.dec.Decode(&msg); nil != err {
		log.Printf("failed to decode message; err: %v", err)
		return msg, err
	}
	return msg, nil
}

func ConnHandler(conn net.Conn) {
	mc := MyConnection{
		conn:       conn,
		recvBuffer: make([]byte, 4096),
	}
	mc.dec = gob.NewDecoder(&mc.codecBuffer)

	for {
		msg, err := mc.RecvMessage()
		if nil != err {
			if io.EOF == err {
				log.Printf("connection is closed from client; %v", mc.conn.RemoteAddr().String())
				return
			}
			log.Printf("failed to recv message! err: %v", err)
			continue
		}

		fmt.Println("msg: ", msg)
	}
}
