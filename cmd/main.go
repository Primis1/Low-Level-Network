package main

import (
	"low-level-tools/cmd/internal/config"
	protocols "low-level-tools/cmd/pkg/protocols/LLHttp"
)

func init() {
	config.SetKeyENV()
}

func main() {
	config.SetKeyENV()	

	protocols.TCPHttpReq()

}
