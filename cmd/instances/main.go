package main

import (
	"low-level-tools/cmd/internal/config"
)

// we initialize EnvVariable 
func init() {
	config.SetKeyENV()
}