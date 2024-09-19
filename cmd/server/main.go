package main

import (
	"flag"
	"log"
	"net"
	"server/pkg/echoUpper"
)

func main() {
	const name = "tcpuppering"
	log.SetPrefix(name + "\t")

	port := flag.Int("p", 8080, "port to listen")
	flag.Parse()

	// Listener create a TCP listener, which accepts all connection to given port 
	// TCPAddr represents the address of TCP endpoint, which has IP and PORT
	listerner, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port}) 
	if err != nil {
		panic(err)
	}

	defer listerner.Close()

	log.Printf("listen at localhost %s", listerner.Addr())
	for { // infinite loop, so we can all-day long 
		  // listen to incoming requests 
		conn, err := listerner.Accept()

		if err != nil {
			panic(err)
		}
		go echoUpper.Echo(conn, conn)
		
	}
}
