package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	username, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", username.Username)
	repl.Start(os.Stdin, os.Stdout)
}
