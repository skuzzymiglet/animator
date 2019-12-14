package main

import (
	"fmt"
	"strconv"
)

func Render(mainExpr string, otherFrames map[string][]float64) string {
	otherExpr := ""

	for k := range otherFrames {
		otherExpr += string(fmt.Sprintf("-i %v ", k))
	}

	otherFilter := ""

	prev := "0"
	count := 0
	for _, v := range otherFrames {
		count++
		otherFilter += string(fmt.Sprintf("[%v][%v] overlay=0:0:enable='between(t,%v,%v)'", prev, count, v[0], v[1]))
		prev = "v" + strconv.Itoa(count)
		otherFilter += fmt.Sprintf("[%v]", prev)
		if count != len(otherFrames) {
			otherFilter += ","
		}
	}

	return fmt.Sprintf("ffmpeg -y -f image2 -pattern_type sequence -i %q %v -filter_complex %q -map \"[%v]\" x.mp4", mainExpr, otherExpr, otherFilter, prev)
}
