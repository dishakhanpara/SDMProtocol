package main

import (
	entity "SAPProto/transport/Entity"
	parser "SAPProto/transport/parsers"
	"net/netip"
	"os"
)

func main() {

	// go clientserver.RunServer("localhost:2020")

	// time.Sleep(1 * time.Second)

	// clientserver.RunClient("localhost:5333", "localhost:2020", "hello")

	// time.Sleep(5 * time.Second)

	laddr, err1 := netip.ParseAddr("233.233.233.233:2020")
	if err1 != nil {
		println(err1.Error())
		os.Exit(-1)
	}

	raddr, err2 := netip.ParseAddr("255.255.255.255:5050")
	if err2 != nil {
		println(err2.Error())
		os.Exit(-1)
	}

	p := entity.NewPackage(1, laddr, raddr, 5, "hello")

	list := parser.ParseToBinary(&p)

	for _, i := range list {
		println(i)
	}

}
