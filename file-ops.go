package main

import (
	"io/ioutil"
	"log"
	//"os"
	"path/filepath"
)

func FrameNames(timeline [][]string, tmp string) [][]string {
	// make all paths absolute
	for i, line := range timeline {
		for j, frame := range line {
			timeline[i][j] = filepath.Join(string(tmp), string(frame))
		}
	}
	return timeline
}

func FileOps(timeline [][]string) ([][]string, string, string) {
	dir, err := ioutil.TempDir("", "animator")
	if err != nil {
		log.Fatal(err)
	}

	timeline = FrameNames(timeline, dir)

	return timeline, dir, ""
}
