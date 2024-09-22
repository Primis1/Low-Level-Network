package config

import (
	"os"
)

func SetKeyENV() {
	os.Setenv("KEY_WORD", "cmd")
}
