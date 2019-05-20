package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
	"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {

	serverAddress := "localhost:1234"

	_, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil {
		log.Fatal("format error", err)
	}

	conn, err := net.DialTimeout("tcp", serverAddress, time.Duration(1) * time.Millisecond)
	if err != nil {
		log.Fatal("new conn fail, addr %s, err %v", serverAddress, err)
	}

	args := &Args{7,8}
	var reply int
	//var reply Quotient

	cli := rpc.NewClient(conn)
	err = cli.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("call fail, err %v", err)
	}

	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)


}
