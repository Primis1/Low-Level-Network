package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	const name = "writetcp"
	log.SetPrefix(name + "\t")

	port := flag.Int("p", 8080, "port to connect")
	flag.Parse()

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: *port})
	if err != nil {
		log.Fatalf("error, connection to localHost: %d,  %v", *port, err)
	}
	log.Printf("connected to %s: will forward stdin", conn.RemoteAddr())

	defer conn.Close()

	// we do read incoming line from the server
	go func() {
		for connScanner := bufio.NewScanner(conn); connScanner.Scan(); { 
			fmt.Printf("%s\n", connScanner.Text())

			if err := connScanner.Err(); err != nil {
				log.Fatal("error reading from %s: %v", conn.RemoteAddr(),  err)
			}
			if connScanner.Err() != nil {
				log.Fatal("error reading from %s: %v", conn.RemoteAddr(),  err)
			}
		}
	}()

	// we do read messages from standart input(terminal), 
	// log if something wrong and send them to server
	for stdinScanner := bufio.NewScanner(os.Stdin); stdinScanner.Scan(); {
		log.Printf("sentL %s\n", stdinScanner.Text())

		if _, err := conn.Write(stdinScanner.Bytes()); err != nil {
			log.Fatal("error occured while writting from %s: %v", conn.RemoteAddr(), err )
		}
		if _, err := conn.Write([]byte("\n")); err != nil {
			log.Fatal("error occured while writting from %s: %v", conn.RemoteAddr(), err )
		}
		if stdinScanner.Err() != nil {
			log.Fatal("error occured while reading from %s: %v", conn.RemoteAddr(), err )
		}
	}
}