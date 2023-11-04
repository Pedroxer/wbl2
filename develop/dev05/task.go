package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"testing"
)

var after string
var before string
var context string
var count bool
var ignore bool
var invert bool
var fixed bool
var line_num bool

func init() {
	testing.Init()
	flag.StringVar(&after, "A", "0", "print n lines after match")   //+
	flag.StringVar(&before, "B", "0", "print n lines before match") //+
	flag.StringVar(&context, "C", "0", "print n lines around match")
	flag.BoolVar(&count, "c", false, "count amount of lines")                         //+
	flag.BoolVar(&ignore, "i", false, "ignore case")                                  //+
	flag.BoolVar(&invert, "v", false, "instead of matching, exclude")                 //+
	flag.BoolVar(&fixed, "F", false, "the exact match with string, not with pattern") //+
	flag.BoolVar(&line_num, "n", false, "print line number")                          //+

	flag.Parse()
}

// FindPattern finds a pattern of reg expr from os.Args.
func FindPattern(args []string) (string, error) {
	for i := range args {
		res, err := regexp.Match(`\-\D+`, []byte(args[i]))
		if err != nil {
			return "", err
		}
		if !res {
			return args[i], err
		}
	}
	return "", nil
}

// readStringsFromFile reads strings from files
func readStringsFromFile(file string) (lines []string, err error) {
	var scanner *bufio.Scanner
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	scanner = bufio.NewScanner(f)
	defer f.Close()
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func buildRegExp(pattern string) (*regexp.Regexp, error) {
	if fixed {
		pattern = `\Q` + pattern + `\E`

	}
	if ignore {
		pattern = `(?i)` + `(` + pattern + `)`
	}
	if invert {

	}
	regExp, err := regexp.Compile(pattern)
	if err != nil {
		return regExp, err
	}
	return regExp, nil

}

// execRegExp execute regexp for array of strings. And check line_num and invert flags
func execRegExp(re regexp.Regexp, lines []string) []string {
	var strs []string
	if line_num {
		for i, val := range lines {
			res := re.Match([]byte(val))
			if invert {
				res = !res
			}
			if res {
				if context != "0" && after == "0" && before == "0" {
					after = context
					before = context
					strs = append(strs, printWhenBefore(&re, lines, i)...)
					strs = append(strs, fmt.Sprintf("Found in line %d: %s", i+1, val))
					strs = append(strs, printWhenAfter(&re, lines, i)...)
				} else {
					if before != "0" {
						strs = append(strs, printWhenBefore(&re, lines, i)...)
					}
					strs = append(strs, fmt.Sprintf("Found in line %d: %s", i+1, val))
					if after != "0" {
						strs = append(strs, printWhenAfter(&re, lines, i)...)
					}
				}
			}
		}
		return strs
	}
	for i, val := range lines {
		res := re.Match([]byte(val))
		if invert {
			res = !res
		}
		if res {
			if context != "0" && after == "0" && before == "0" {
				after = context
				before = context
				strs = append(strs, printWhenBefore(&re, lines, i)...)
				strs = append(strs, val)
				strs = append(strs, printWhenAfter(&re, lines, i)...)
			} else {
				if before != "0" {
					strs = append(strs, printWhenBefore(&re, lines, i)...)
				}
				strs = append(strs, val)
				if after != "0" {
					strs = append(strs, printWhenAfter(&re, lines, i)...)
				}
			}
		}
	}
	return strs
}

func maxInt(a, b int) string {
	if a > b {
		return fmt.Sprint(a)
	}
	return fmt.Sprint(b)
}

func printWhenAfter(re *regexp.Regexp, lines []string, core_pos int) []string {
	var strs []string
	after_num, err := strconv.Atoi(after)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if line_num {
		for j := core_pos + 1; j <= core_pos+after_num; j++ {
			if re.Match([]byte(lines[j])) {
				break
			}
			strs = append(strs, fmt.Sprintf("Line %d: %s", j+1, lines[j]))
		}
	} else {
		for j := core_pos + 1; j <= core_pos+after_num; j++ {
			if re.Match([]byte(lines[j])) {
				break
			}
			strs = append(strs, lines[j])
		}
	}
	strs = append(strs, "--")
	return strs
}

func printWhenBefore(re *regexp.Regexp, lines []string, core_pos int) []string {
	var strs []string
	strs = append(strs, "--")
	before_num, err := strconv.Atoi(before)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if line_num {
		for j := core_pos - before_num; j < core_pos; j++ {
			if re.Match([]byte(lines[j])) {
				break
			}
			strs = append(strs, fmt.Sprintf("Line %d: %s", j+1, lines[j]))
		}
	} else {
		for j := core_pos - before_num; j < core_pos; j++ {
			if re.Match([]byte(lines[j])) {
				break
			}

			strs = append(strs, lines[j])
		}
	}
	return strs
}

func main() {
	filePath := os.Args[len(os.Args)-1]
	strs, err := readStringsFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}
	pattern, _ := FindPattern(os.Args)
	re, err := buildRegExp(pattern)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	out := execRegExp(*re, strs)
	if count {
		fmt.Println(len(out))
		return
	}
	for _, val := range out {
		fmt.Println(val)
	}

}
