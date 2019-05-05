package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"

	ms "github.com/multiformats/go-multistream"
)

// This example creates a multistream muxer, adds handlers for the protocols
// "/cats" and "/dogs" and exposes it on a localhost:8765. It then opens connections
// to that port, selects the protocols and tests that the handlers are working.
func main() {
	mux := ms.NewMultistreamMuxer()
	mux.AddHandler("/cats", func(proto string, rwc io.ReadWriteCloser) error {
		fmt.Fprintln(rwc, proto, ": HELLO I LIKE CATS")
		return rwc.Close()
	})
	mux.AddHandler("/dogs", func(proto string, rwc io.ReadWriteCloser) error {
		fmt.Fprintln(rwc, proto, ": HELLO I LIKE DOGS")
		return rwc.Close()
	})

	list, err := net.Listen("tcp", ":8765")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			con, err := list.Accept()
			if err != nil {
				panic(err)
			}

			go mux.Handle(con)
		}
	}()

	// The Muxer is ready, let's test it
	conn, err := net.Dial("tcp", ":8765")
	if err != nil {
		panic(err)
	}

	// Create a new multistream to talk to the muxer
	// which will negotiate that we want to talk with /cats
	mstream := ms.NewMSSelect(conn, "/cats")
	cats, err := ioutil.ReadAll(mstream)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", cats)
	mstream.Close()

	// A different way of talking to the muxer
	// is to manually selecting the protocol ourselves
	conn, err = net.Dial("tcp", ":8765")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = ms.SelectProtoOrFail("/dogs", conn)
	if err != nil {
		panic(err)
	}
	dogs, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", dogs)
	conn.Close()
}
