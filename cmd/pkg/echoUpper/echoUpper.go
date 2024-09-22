package echoUpper

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func Echo(w io.Writer ,r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		
		fmt.Fprintf(w, "%s\n", strings.ToUpper(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error, reading from %s", err)
	}
}