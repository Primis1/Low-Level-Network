package config

import "os"

//
// for trimming the folders path in logging, i have decided to
// put the keyword of main folder into ENV, just for case someone decides
// to use a different folder structure
//

func init() {
	os.Setenv("KEY_WORD", "cmd")
}