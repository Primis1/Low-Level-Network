package echoUpper

import (
	"bufio"
	"io"
	"low-level-tools/cmd/pkg/logging"
	"strings"
)

func Echo(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)
	log := logging.NewLogger(logging.INFO)
	errMsg := logging.NewLogger(logging.ERR)

	for scanner.Scan() {
		line := scanner.Text()

		log.Info(w, "%s\n", strings.ToUpper(line))
	}

	if err := scanner.Err(); err != nil {
		errMsg.Error("Error, reading from %s", err)
	}
}
