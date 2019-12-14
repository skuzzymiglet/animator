package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	re     string
	expand func([]byte) []byte
}

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

func expandRepeat(inText []byte) []byte {
	s := ""
	text := string(inText)
	timesString := regexp.MustCompile("(\\d+)\\(.+\\)").FindStringSubmatch(text)[1]
	times, _ := strconv.Atoi(timesString)
	if times == 0 {
		return inText
	}
	inside := regexp.MustCompile("\\d+\\((.+)\\)").FindStringSubmatch(text)[1]
	for i := 0; i != times; i += 1 {
		s += inside
		if i != times-1 {
			s += ","
		}
	}
	return []byte(s)
}

func ReplaceAll(s string) string {
	rules := []rule{rule{"[0-9]+\\-[0-9]+", expandRange}, rule{"\\d+\\([^\\)]+\\)", expandRepeat}}
	for r := range rules {
		s = string(regexp.MustCompile(rules[r].re).ReplaceAllFunc([]byte(s), rules[r].expand))
	}
	return s
}
