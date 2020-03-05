package main

import (
	"./repl"
)

func main() {
	repl := repl.Make(64)
	repl.Run()
}
