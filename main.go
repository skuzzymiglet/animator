package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func testParse() {
	fmt.Println("Parse Test")
	fmt.Println(ReplaceAll("8(5-10,7),7(7-10)"))
	fmt.Println(ReplaceAll("[3(7-9)][4,2(6-10)]"))
	fmt.Println(Files("1,2,3", []string{"x.jpg", "y.png", "z.webm", "f.d"}))
	fmt.Println(Files(ReplaceAll("[3(1-3)][4,2(2-4)]"), []string{"hi.png", "me.png", "me.webm", "yeet.xcf"}))
}

func testRender() {
	fmt.Println("Render Test")
	main := "i%2d.jpg"
	other := map[string][]float64{}
	other["o0.jpg"] = []float64{0, 0.1}
	other["o1.jpg"] = []float64{0.1, 0.4}
	other["o3.webm"] = []float64{1, 3}
	fmt.Println(Render(main, other, "hello.webm"))
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
