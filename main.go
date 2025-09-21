package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Printf("🙈 🙉 🙊\n")
	repl.Start(os.Stdin, os.Stdout)
}
