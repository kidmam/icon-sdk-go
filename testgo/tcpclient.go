package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"net"
	"time"
)

func init() {
	gob.Register(MyMsgBodyPing{})
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
	enc         *gob.Encoder
	codecBuffer bytes.Buffer
}

func (mc *MyConnection) SendMessage(msg MyMsg) error {
	mc.codecBuffer.Reset()

	mc.enc.Encode(msg)

	lengthBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthBuf, uint32(mc.codecBuffer.Len()))

	if _, err := mc.conn.Write(lengthBuf); nil != err {
		log.Printf("failed to send msg length; err: %v", err)
		return err
	}

	if _, err := mc.conn.Write(mc.codecBuffer.Bytes()); nil != err {
		log.Printf("failed to send msg; err: %v", err)
		return err
	}

	return nil
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5032")
	if nil != err {
		log.Fatalf("failed to connect to server")
	}

	mc := MyConnection{
		conn: conn,
	}
	mc.enc = gob.NewEncoder(&mc.codecBuffer)

	for {
		mc.SendMessage(MyMsg{
			Header: MyMsgHeader{
				MsgType: "ping",
				Date:    time.Now().UTC().Format(time.RFC3339),
			},
			Body: MyMsgBodyPing{
				Content: "Hello! I'm alive!",
			},
		})

		time.Sleep(time.Duration(3) * time.Second)
	}
}
