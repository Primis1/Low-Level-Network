package protocols

import (
	"bufio"
	"flag"
	"fmt"
	"low-level-tools/cmd/pkg/logging"
	"net"
	"os"
	"strings"
)

var (
	host, url, method string
	port              int
)

func TCPHttpReq() {
	//
	info := logging.NewLogger(logging.INFO)
	errorMsg := logging.NewLogger(logging.ERR)
	//

	//
	flag.StringVar(&method, "method", "GET", "method of HTTP request")
	flag.StringVar(&host, "host", "localhost", "host of resource")
	flag.StringVar(&url, "path", "/", "path to resource")
	flag.IntVar(&port, "port", 8080, "port to resource")

	flag.Parse()
	//

	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		errorMsg.Error(err)
	}

	conn, err := net.DialTCP("tcp", nil, ip)
	if err != nil {
		errorMsg.Error(err)
	}

	info.Info("connected to %s: %s", host, conn.RemoteAddr())

	defer conn.Close()

	var reqFields = []string{
		fmt.Sprintf("%s %s HTTP/1.1", method, host),
		"Host: " + host,
		"User-agent: httpget",
		"\n", 
	}

	request := strings.Join(reqFields, "\r\n") + "\r\n"

	conn.Write([]byte(request))
	info.Info("we just wrote a request")

	for scanner := bufio.NewScanner(conn); scanner.Scan(); {
		line := scanner.Bytes()
		if _, err := fmt.Fprintf(os.Stdout, "%s", line); err != nil {
			errorMsg.Error("scanning writting went wrong\t", err)
		}
		if scanner.Err() != nil {
			errorMsg.Error("reading went wrong\t", err)
			return
		}

	}
}