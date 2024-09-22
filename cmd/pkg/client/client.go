package client

import (
	"bufio"
	"flag"
	"fmt"
	"low-level-tools/cmd/pkg/logging"
	"net"
	"os"
)

func Client() {
	log := logging.NewLogger(logging.INFO)
	errMsg := logging.NewLogger(logging.ERR)

	const name = "writetcp"

	port := flag.Int("p", 8080, "port to connect")
	flag.Parse()

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: *port})
	if err != nil {
		errMsg.Error("error, connection to localHost: %d,  %v \n", *port, err)
	}
	log.Info("connected to %s: will forward stdin \n", conn.RemoteAddr())

	defer conn.Close()

	// we do read incoming line from the server
	go func() {
		for connScanner := bufio.NewScanner(conn); connScanner.Scan(); {
			fmt.Printf("%s\n", connScanner.Text())

			if err := connScanner.Err(); err != nil {
				errMsg.Error("error reading from %s: %v", conn.RemoteAddr(), err)
			}
			if connScanner.Err() != nil {
				errMsg.Error("error reading from %s: %v", conn.RemoteAddr(), err)
			}
		}
	}()

	// we do read messages from standart input(terminal),
	// log if something wrong and send them to server
	for stdinScanner := bufio.NewScanner(os.Stdin); stdinScanner.Scan(); {
		log.Info("sent %s\n", stdinScanner.Text())

		if _, err := conn.Write(stdinScanner.Bytes()); err != nil {
			errMsg.Error("error occured while writting from %s: %v", conn.RemoteAddr(), err)
		}
		if _, err := conn.Write([]byte("\n")); err != nil {
			errMsg.Error("error occured while writting from %s: %v", conn.RemoteAddr(), err)
		}
		if stdinScanner.Err() != nil {
			errMsg.Error("error occured while reading from %s: %v", conn.RemoteAddr(), err)
		}
	}
}
