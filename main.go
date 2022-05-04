package main

import (
	"os"

	repl "github.com/ZooeyLang/REPL"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
