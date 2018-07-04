package main

import (
	"net/rpc"
	"log"
	"fmt"
	"server"
)

func main() {
	client, err := rpc.DialHTTP("tcp",  ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &server.Args{7,8}
	var reply int
	call := client.Go("Arith.Multiply", args, &reply, nil)
	replyCall := <-call.Done    // will be equal to divCall
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, *(replyCall.Reply.(*int)))
}
