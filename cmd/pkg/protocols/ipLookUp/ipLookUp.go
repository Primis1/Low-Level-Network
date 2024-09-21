package protocols

import (
	"net"
	"os"
	"low-level-tools/cmd/models/logging"
)

func IpLookUp() {
	log := logging.NewLogger(logging.INFO)
	errMsg := logging.NewLogger(logging.ERR)

	logging.Some()
	if len(os.Args) != 2  {
		log.Info("Should use DNS of the website to get IP")
		errMsg.Error("expected exactly one argument; got %d", len(os.Args)-1)
	}

	host := os.Args[1]
	
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Info("lookup ip: %s: %v", host, err)
	}

	if len(ips) == 0 {
		errMsg.Error("no ips found for %s", host)
	}

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
