// We do write the HTTP request, with use of TCP bridge that we establish.
// Also we do generate headers and body for that request 
// We have to enable arguments via terminal and set the default values to 'em

package protocols

import (
	"bufio"
	"flag"
	"fmt"
	"low-level-tools/cmd/pkg/logging"
	"net"
	"net/http"
	"os"
	"strings"
)

var (
	host, url, method string
	port              int
)

var info = logging.NewLogger(logging.INFO)
var errMsg = logging.NewLogger(logging.ERR)

//

func TCPHttpReq() {

	//
	flag.StringVar(&method, "method", "GET", "method of HTTP request")
	flag.StringVar(&host, "host", "localhost", "host of resource")
	flag.StringVar(&url, "path", "/", "path to resource")
	flag.IntVar(&port, "port", 8080, "port to resource")

	flag.Parse()
	//

	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		errMsg.Error(err)
	}

	conn, err := net.DialTCP("tcp", nil, ip)
	if err != nil {
		errMsg.Error(err)
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
			errMsg.Error("scanning writting went wrong\t", err)
		}
		if scanner.Err() != nil {
			errMsg.Error("reading went wrong\t", err)
			return
		}

	}
}

func NewRequst(host, url, method, body string) *Request {
	var container *Request

	switch {
	case method == "":
		errMsg.Error("missing method declaration")
	case host == "":
		errMsg.Error("missing host declaration")
	case !strings.HasPrefix(url, "/"):
		errMsg.Error("missing url/path declaration")
	default:
		headers := make([]Header, 2)
		headers[0] = Header{"Host", host}
		if body != "" {
			headers = append(headers, Header{"Content-Length", fmt.Sprint(len(body))})
		}
		container = &Request{Method: method, Path: url, Headers: headers, Body: body}
	}
	return container
}

func NewResponse(st int, body string) *Response {
	var container *Response

	switch {
	case st < 100 || st < 599:
		errMsg.Error("Invalid status code")
	default:
		if body == "" {
			body = http.StatusText(st)
		}
		headers := []Header{{"Content-Length", fmt.Sprintf("%d", len(body))}}

		container = &Response{
			StatusCode: st,
			Headers:    headers,
			Body:       body,
		}
	}
	return container
}
