package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	// Replacement Rule
	re     string              // What to search for
	expand func(string) string // What to do on it (replaceAllFunc)
}

func expandRange(text string) string {
	// Expand range e.g. 4-7, into comma separated e.g 4,5,6,7
	s := ""
	start, serr := strconv.Atoi(strings.Split(text, "-")[0]) // 4-7 -> 4
	end, eerr := strconv.Atoi(strings.Split(text, "-")[1])   // 4-7 -> 7
	if serr != nil || eerr != nil {
		fmt.Printf("err: %q\n", text)
	}
	// RANGE
	counter := 0
	if end-start >= 0 {
		counter = 1
	} else {
		counter = -1
	}
	for i := start; i != end+counter; i += counter {
		// TODO factor
		if i == end {
			s += strconv.Itoa(i) // no comma at end
		} else {
			s += strconv.Itoa(i) + ","
		}
	}
	return s
}

func expandRepeat(text string) string {
	// Expand repeat e.g 4(3), into comma separated e.g 3,3,3,3
	s := ""
	timesString := regexp.MustCompile("(\\d+)\\(.+\\)").FindStringSubmatch(text)[1] // 6(7-90) -> 6
	times, _ := strconv.Atoi(timesString)
	// do not do 0 times (repeat till end)
	// TODO expandRepeatTillEnd
	if times == 0 {
		return text
	}

	inside := regexp.MustCompile("\\d+\\((.+)\\)").FindStringSubmatch(text)[1] // 6(7-90) -> 7-90
	// append inside to returning string `times` times
	for i := 0; i != times; i++ {
		s += inside
		if i != times-1 {
			s += ","
		}
	}
	return s
}

func ReplaceAll(s string) string {
	// Replace (expand?) all according to a list of `rule`s
	rules := []rule{
		rule{"[0-9]+\\-[0-9]+", expandRange},
		rule{"\\d+\\([^\\)]+\\)", expandRepeat},
	}
	for r := range rules {
		s = string(regexp.MustCompile(rules[r].re).ReplaceAllStringFunc(s, rules[r].expand))
	}
	return s
}

func Files(s string, files []string) string {
	re := regexp.MustCompile("\\d+")
	s = re.ReplaceAllStringFunc(s, func(index string) string {
		i, _ := strconv.Atoi(index)
		return fmt.Sprintf("%q", files[i-1])
	})
	return s
}
