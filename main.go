package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	name := "animator"
	parser := argparse.NewParser(name, "command-line animation toolkit")
	expr := parser.String("e", "expr", &argparse.Options{Required: true, Help: "Expression"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	fmt.Println(*expr)
}
