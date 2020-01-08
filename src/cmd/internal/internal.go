package internal

import (
	"bufio"
	"fmt"
	"os"
)

var homedir string

func init() {
	homedir, _ = os.UserHomeDir()
}

func InstallStaticFiles() {

}

func ConfFileExists() bool {
	stat, _ := os.Stat(homedir + "/.envconf")
	return stat != nil
}

func Reconfigure() {
	fmt.Println("Reconfiguring env")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Git name: ")
	var gitUsername string
	if scanner.Scan() {
		gitUsername = scanner.Text()
	} else {
		fmt.Println("Failed to read git username")
		os.Exit(1)
	}

	fmt.Print("Git mail: ")
	var gitEmail string
	if scanner.Scan() {
		gitEmail = scanner.Text()
	}

	fmt.Printf("Github information: %s <%s>", gitUsername, gitEmail)

	conf := Configuration{}
	conf.Git.Name = gitUsername
	conf.Git.Email = gitEmail

	fmt.Printf("%s\n", conf.Marshal())
}

