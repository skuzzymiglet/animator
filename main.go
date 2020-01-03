package main

import (
	"bytes"
	"fmt"
	"github.com/akamensky/argparse"
	"os"
	"os/exec"
	"strings"
)

func testParse() {
	fmt.Println("Parse Test")
	fmt.Println(ReplaceAll("8(5-10,7),7(7-10)"))
	fmt.Println(ReplaceAll("[3(7-9)][4,2(6-10)]"))
	fmt.Println(Files("1,2,3", []string{"x.jpg", "y.png", "z.webm", "f.d"}))
	fmt.Println(Files(ReplaceAll("[3(1-3)][4,2(2-4)]"), []string{"hi.png", "me.png", "me.webm", "yeet.xcf"}))
	fmt.Println(StringToTimeline(Files(ReplaceAll("[3(1-3)][4,2(2-4)]"), []string{"hi.png", "me.png", "me.webm", "yeet.xcf"})))
	fmt.Println("FileOps Test")
	tl, tmp := FileOps(StringToTimeline(Files(ReplaceAll("[3(1-3)][4,2(2-4),_]"), []string{"hi.png", "me.png", "me.webm", "yeet.xcf"})))
	defer os.RemoveAll(tmp)
	fmt.Printf("%#v\n", tl)
	frames := TimelineToFrames(tl, 30)
	fmt.Println(frames)
	fmt.Println(Render(frames, "f.webm"))
}

func testRender() {
	fmt.Println("Render Test")
	//main := "i%2d.jpg"
	other := map[string][]float64{}
	other["o0.jpg"] = []float64{0, 0.1}
	other["o1.jpg"] = []float64{0.1, 0.4}
	other["o3.webm"] = []float64{1, 3}
	fmt.Println(Render(other, "hello.webm"))
}

func main() {
	//testParse()
	//testRender()
	name := "animator"

	parser := argparse.NewParser(name, "command-line animation toolkit")
	expr := parser.String("e", "expr", &argparse.Options{Required: true, Help: "Expression"})
	files := parser.List("f", "files", &argparse.Options{Required: true, Help: "List of input files"})
	output := parser.String("o", "out", &argparse.Options{Required: true, Help: "Output file"})
	fps := parser.Float("r", "rate", &argparse.Options{Required: false, Help: "Framerate", Default: 5.0})

	parseErr := parser.Parse(os.Args)
	if parseErr != nil {
		fmt.Print(parser.Usage(parseErr))
	}
	tl, tmp := FileOps(StringToTimeline(Files(ReplaceAll(*expr), *files)))
	defer os.RemoveAll(tmp)
	fmt.Printf("Timeline: %#v\n", tl)
	frames := TimelineToFrames(tl, *fps)
	fmt.Println("Frames:", frames)
	fields := strings.Split(Render(frames, *output), " ")
	cmd := exec.Command(fields[0], fields[1:]...)
	fmt.Printf("Fields: %#v\n", fields)
	fmt.Println("Command", cmd)

	var out bytes.Buffer
	var err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err

	x := cmd.Run()

	fmt.Println(x, out.String(), err.String())
}
