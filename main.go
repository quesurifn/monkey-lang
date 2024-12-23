package main

import (
	"fmt"
	"os"
	user2 "os/user"
	"quesurifn/monkey/repl"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n\n", user)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
