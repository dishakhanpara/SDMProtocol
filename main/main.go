package main

import (
	clientserver "SAPProto/client-server"
	"time"
)

func main() {

	go clientserver.RunServer("localhost:2020")

	time.Sleep(1 * time.Second)

	clientserver.RunClient("localhost:5333", "localhost:2020", "hello")

	time.Sleep(5 * time.Second)

}
