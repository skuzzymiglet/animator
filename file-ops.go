package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FrameNames(timeline [][]string, tmp string) [][]string {
	// make all paths absolute
	// move coherent topline frames to mx.y
	for i, line := range timeline {
		for j, frame := range line {
			timeline[i][j] = string(tmp) + string(frame)
		}
	}
	return timeline
}

func FileOps(timeline [][]string, files []string) ([][]string, string, string) {
	dir, err := ioutil.TempDir("", "animator")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range files {
		os.Symlink(e, filepath.Join(dir, filepath.Base(e)))
	}

	//ls, err := ioutil.ReadDir(dir)

	return timeline, dir, ""
}
