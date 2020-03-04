package main

import "./repl"

func main() {
	repl := repl.Make(2)
	repl.Run()
}
