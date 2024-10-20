package protocols

import (
	"low-level-tools/cmd/pkg/logging"
	"net"
	"os"
)

// What steps we do:
// initialize the logger, how we can do without it anyways?
// we use standard input, i.e terminal command before running the executable
// 		we set/receive DNS of the resource

func IpLookUp() {

	log := logging.NewLogger(logging.INFO)
	errMsg := logging.NewLogger(logging.ERR)

	if len(os.Args) != 2 {
		log.Info("Should use DNS of the website to get IP")
		errMsg.Error("expected exactly one argument; got %d", len(os.Args)-1)
	}
	host := os.Args[1]

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Error("lookup ip: %s: %v", host, err)
	}

	if len(ips) == 0 {
		errMsg.Error("no ips found for %s", host)
	}
	log.Info(ips)
	for _, ip := range ips {
		if ip.To4() != nil {
			log.Info(ip)
			return
		}
	}
	for _, ip := range ips {
		if ip.To16() != nil {
			log.Info(ip)
			return
		}
	}
}
