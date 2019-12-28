package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FrameNames(timeline [][]string, tmp string) [][]string {
	// make all paths absolute
	var newName string
	for i, line := range timeline {
		for j, frame := range line {
			newName = filepath.Join(string(tmp), string(frame))
			oldNameAbs, err := filepath.Abs(timeline[i][j])
			if err != nil {
				log.Fatal(err)
			}
			os.Symlink(oldNameAbs, newName)
			timeline[i][j] = newName
		}
	}
	return timeline
}

func FileOps(timeline [][]string) [][]string {
	dir, err := ioutil.TempDir("", "animator")
	if err != nil {
		log.Fatal(err)
	}

	timeline = FrameNames(timeline, dir)

	return timeline
}
