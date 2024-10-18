// clientserver package consist of two main function, which we should execute separetly
// one is client which writes a message to the "server" which should return the upper cased string, as a response

// it uses command line flags, and bufio as a line-by-line reader

package clientserver

import (
	"bufio"
	"flag"
	"io"
	"low-level-tools/cmd/pkg/logging"
	"net"
	"strings"
)

func Server() {
	// Define the log and err bridges
	log := logging.NewLogger(logging.INFO)
	errMsg := logging.NewLogger(logging.ERR)

	const name = "tcpuppering"
	log.Info(name + "\t")

	port := flag.Int("p", 8080, "port to listen")
	flag.Parse()

	// Listener create a TCP listener, which accepts all connection to given port
	// TCPAddr represents the address of TCP endpoint, which has IP and PORT
	listerner, err := net.ListenTCP("tcp", &net.TCPAddr{Port: *port})
	if err != nil {
		errMsg.Error(err)
	}

	defer listerner.Close()

	log.Info("listen at localhost %s", listerner.Addr())
	for {
		// infinite loop, so we can all-day long -
		// - listen to incoming requests
		conn, err := listerner.Accept()

		if err != nil {
			panic(err)
		}
		go Echo(conn, conn)

	}
}

func Echo(w io.Writer, r io.Reader) {
	// bufio scanner works as an OS or IO reader in the files, but reads it line-by-line  
	scanner := bufio.NewScanner(r)
	log := logging.NewLogger(logging.INFO)
	errMsg := logging.NewLogger(logging.ERR)

	// run through, untill there is an error or empty 
	for scanner.Scan() {
	// receive last writen string-line  
		line := scanner.Text()

		log.Info(w, "%s\n", strings.ToUpper(line))
	}

	if err := scanner.Err(); err != nil {
		errMsg.Error("Error, reading from %s", err)
	}
}
