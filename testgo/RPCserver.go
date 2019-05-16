package main

import (
	"fmt"
	"github.com/kidmam/icon-sdk-go/testgo/shardRPC"
	"math"
	"net"
	"net/rpc"
	"os"
)

type MyInterface struct{}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (t *MyInterface) Multiply(arguments *shardRPC.MyFloats, reply *float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

func (t *MyInterface) Power(arguments *shardRPC.MyFloats, reply *float64) error {
	*reply = Power(arguments.A1, arguments.A2)
	return nil
}

func main() {
	PORT := ":1234"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	myInterface := new(MyInterface)
	rpc.Register(myInterface)
	t, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
