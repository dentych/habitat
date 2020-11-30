package terminal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func Clear() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func Read() string {
	fmt.Print("Input: ")
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ReadByte() byte {
	fmt.Print("Input: ")
	if scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			return text[0]
		}
	}
	return 0
}

func ReadEnter() {
	if scanner.Scan() {
		return
	}
}