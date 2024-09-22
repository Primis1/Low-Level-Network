package main

import (
	"low-level-tools/cmd/config"
	protocols "low-level-tools/cmd/pkg/protocols/LLHttp"
)

func main() {
	config.SetKeyENV()

	protocols.TCPHttpReq()
}
