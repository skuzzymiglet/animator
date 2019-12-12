package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func expandRange(inText []byte) []byte {
	s := ""
	text := string(inText)
	start, serr := strconv.Atoi(strings.Split(text, "-")[0])
	end, eerr := strconv.Atoi(strings.Split(text, "-")[1])
	if serr != nil || eerr != nil {
		fmt.Printf("err: %q\n", text)
	}
	counter := 0
	if end-start >= 0 {
		counter = 1
	} else {
		counter = -1
	}
	for i := start; i != end+counter; i += counter {
		if i == end {
			s += strconv.Itoa(i)
		} else {
			s += strconv.Itoa(i) + ","
		}
	}
	return []byte(s)
}

func replaceRange(s string) string {
	search := regexp.MustCompile("[0-9]+\\-[0-9]+")
	return string(search.ReplaceAllFunc([]byte(s), expandRange))
}
