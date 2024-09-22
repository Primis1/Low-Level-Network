package main

import (
	"low-level-tools/cmd/config"
	"low-level-tools/cmd/pkg/logging"
)

func main() {
	config.SetKeyENV()

	logging.NewLogger(logging.INFO).Info("something")
}
