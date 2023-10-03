package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/matheusgb/marmota/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Print("Welcome to Rinha programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
