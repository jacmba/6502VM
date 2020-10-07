package main

import (
	"6502VM/repl"
)

func main() {
	repl := repl.Make(64)
	repl.Run()
}
