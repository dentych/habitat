package internal

import (
	"os"
)

var HomeDir string

func init() {
	HomeDir, _ = os.UserHomeDir()
}
