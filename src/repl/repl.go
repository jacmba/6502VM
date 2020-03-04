package repl

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"../cpu"
	"../mem"
)

/*
REPL Type to run VM in command line interactive mode
*/
type REPL struct {
	mem *mem.Mem // Memory reference
	cpu *cpu.CPU // VM CPU
}

/*
Make creates a REPL instance
*/
func Make(memsize uint16) *REPL {
	mem := mem.Make(memsize)
	cpu := cpu.Make(mem)

	return &REPL{
		mem,
		cpu,
	}
}

/*
Run launches REPL execution
*/
func (repl *REPL) Run() {
	fmt.Println("Starting VM REPL...")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("# ")
		line, _, err := reader.ReadLine()

		if err != nil {
			panic(err)
		}

		input := string(line)

		switch input {
		case "memdump":
			fmt.Println("Printing VM memory dump")
			fmt.Printf("%+v\n", repl.mem.Dump())
			fmt.Println("End of VM memory dump")
		case "cpudump":
			fmt.Println("Printing CPU dump")
			fmt.Printf("%+v\n", repl.cpu)
			fmt.Println("End of CPU dump")
		case "quit":
			println("Bye!")
			os.Exit(0)
		default:
			inputBytes := strings.Split(input, " ")
			bytes, err := hex.DecodeString(strings.Join(inputBytes, ""))
			if err != nil {
				panic(err)
			}

			for i, b := range bytes {
				repl.mem.SetByte(uint16(i), b)
			}

			repl.cpu.RunOnce()
		}
	}
}
