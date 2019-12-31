package main

import (
	"fmt"
	"strconv"
)

// TODO The map[string][]float64 structure used for frames and in-out times removes re-used frames, and must be replaced with a list of structs

func TimelineToFrames(timeline [][]string, fps float64) map[string][]float64 {
	frames := map[string][]float64{}
	for _, x := range timeline {
		for j, y := range x {
			frames[y] = []float64{float64(j) / fps, (float64(j) / fps) + (1.0 / fps)}
		}
	}
	return frames
}

func Render(otherFrames map[string][]float64, out string) string {
	// Take a file pattern and a list of filenames and in-out times and return an `ffmpeg` command to turn them into an output file
	otherExpr := ""

	for k := range otherFrames {
		otherExpr += string(fmt.Sprintf("-i %v ", k)) // Input all otherFrames
	}

	otherFilter := ""

	prev := "0"
	count := 0
	for _, v := range otherFrames {
		// Add text to `filter_complex` overlay
		otherFilter += string(fmt.Sprintf("[%v][%v]overlay=0:0:enable='between(t,%v,%v)'", prev, count, v[0], v[1]))
		prev = "v" + strconv.Itoa(count)
		otherFilter += fmt.Sprintf("[%v]", prev)
		count++
		if count != len(otherFrames) {
			otherFilter += "," // No comma at end
		}
	}
	// format `ffmpeg` command
	return fmt.Sprintf("ffmpeg -y -f image2 -pattern_type sequence %v -filter_complex %v -map [%v] %v", otherExpr, otherFilter, prev, out)
}
