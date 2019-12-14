package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func testParse() {
	fmt.Println("Parse Test")
	fmt.Println(ReplaceAll("8(5-10,7),7(7-10)"))
}

func testRender() {
	fmt.Println("Render Test")
	main := "i%2d.jpg"
	other := map[string][]float64{}
	other["o0.jpg"] = []float64{0, 0.1}
	other["o1.jpg"] = []float64{0.1, 0.4}
	other["o3.webm"] = []float64{1, 3}
	fmt.Println(Render(main, other))
}

func main() {
	testParse()
	testRender()
	name := "animator"
	parser := argparse.NewParser(name, "command-line animation toolkit")
	expr := parser.String("e", "expr", &argparse.Options{Required: true, Help: "Expression"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	fmt.Println(ReplaceAll(*expr))
}
