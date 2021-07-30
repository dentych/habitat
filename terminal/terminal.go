package terminal

import (
	"bufio"
	"bytes"
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

func PrintHeading(heading string) {
	maxLength := 60
	if len(heading) >= maxLength {
		fmt.Println(heading)
		return
	}

	dashAmount := (maxLength - len(heading) - 2) / 2
	var buf bytes.Buffer
	for i := 0; i < dashAmount; i++ {
		buf.WriteRune('-')
	}
	buf.WriteRune(' ')
	buf.WriteString(heading)
	buf.WriteRune(' ')
	for i := 0; i < dashAmount; i++ {
		buf.WriteRune('-')
	}
	for buf.Len() < maxLength {
		buf.WriteRune('-')
	}

	fmt.Println(buf.String())
}
