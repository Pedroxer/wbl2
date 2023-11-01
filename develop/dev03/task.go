package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var key string
var numeric bool
var reverse bool
var unique bool

func init() {
	testing.Init()
	flag.StringVar(&key, "k", "0", `sort with a k's column`)
	flag.BoolVar(&numeric, "n", false, "compare according to string numerical value")
	flag.BoolVar(&reverse, "r", false, "reverse the result of comparisons")
	flag.BoolVar(&unique, "u", false, "output only unique strings")

	flag.Parse()
}

// имплементируем sort.Interface для лексикографической сортировки
type stringSort struct {
	sl  [][]string
	key int
}

func (s stringSort) Len() int {
	return len(s.sl)
}
func (s stringSort) Swap(i, j int) {
	s.sl[i], s.sl[j] = s.sl[j], s.sl[i]
}
func (s stringSort) Less(i, j int) bool {
	jLen := len(s.sl[j])
	iLen := len(s.sl[i])
	jHasKey := iLen > s.key
	iHasKey := jLen > s.key

	switch {
	case jHasKey && iHasKey:
		if s.sl[i][s.key] < s.sl[j][s.key] {
			return true
		}
		if s.sl[i][s.key] > s.sl[j][s.key] {
			return false
		}
	case !iHasKey && jHasKey:
		return true
	case iHasKey && !jHasKey:
		return false
	}

	for k := 0; k < minInt(iLen, jLen); k++ {
		if s.sl[i][k] < s.sl[j][k] {
			return true
		}
		if s.sl[i][k] > s.sl[j][k] {
			return false
		}
	}
	return iLen < jLen
}

// имплементация sort.Interface для сортировки чисел
type numSort struct {
	sl  [][]string
	key int
}

func (s numSort) Len() int {
	return len(s.sl)
}
func (s numSort) Swap(i, j int) {
	s.sl[i], s.sl[j] = s.sl[j], s.sl[i]
}

func (s numSort) Less(i, j int) bool {
	iLen := len(s.sl[i])
	jLen := len(s.sl[j])
	iHasKey := iLen > s.key
	jHasKey := jLen > s.key

	switch {
	case iHasKey && jHasKey:
		if lessNum(s.sl[i][s.key], s.sl[j][s.key]) {
			return true
		}
		if lessNum(s.sl[j][s.key], s.sl[i][s.key]) {
			return false
		}
	case !iHasKey && jHasKey:
		return true
	case iHasKey && !jHasKey:
		return false
	}

	for k := 0; k < minInt(iLen, jLen); k++ {
		if lessNum(s.sl[i][k], s.sl[j][k]) {
			return true
		}
		if lessNum(s.sl[j][k], s.sl[i][k]) {
			return false
		}
	}

	return iLen < jLen
}

// Функция lessNum сравнивает строки как числа и возвращает true,
// если i < j либо j является числом, а i - нет.
func lessNum(i, j string) bool {
	iFloat, iIsNum := stringToFloat(i)
	jFloat, jIsNum := stringToFloat(j)
	if jIsNum && (!iIsNum || iFloat < jFloat) {
		return true
	}
	return false
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Функция stringToFloat конвертирует строку в число и возвращает его,
// а также идикатор успешного выполнения.
// Причем stringToFloat("NaN") = 0, false.
func stringToFloat(s string) (float64, bool) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil || math.IsNaN(f) {
		return 0, false
	}
	return f, true
}

func readStrings(file string) (lines []string, err error) {
	var scanner *bufio.Scanner
	if file == "" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		scanner = bufio.NewScanner(f)
		defer f.Close()
	}

	if unique { // если указан unique, то мы создаём сет, чтобы исключить повторения
		uLines := make(map[string]struct{})
		for scanner.Scan() {
			line := scanner.Text()
			if _, ok := uLines[line]; !ok { // смотрим, есть ли строка в сете
				uLines[line] = struct{}{}
				lines = append(lines, line)
			}
		}
	} else {
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}
	return lines, nil
}

func readStringSlices(file string, k int) (lines [][]string, err error) {
	var scanner *bufio.Scanner
	if file == "" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		scanner = bufio.NewScanner(f)
		defer f.Close()
	}
	if unique {
		uLines := make(map[string]struct{})
		for scanner.Scan() {
			line := strings.Fields(scanner.Text())
			var field string
			if k < len(line) {
				field = line[k]
			}
			if _, ok := uLines[field]; !ok {
				uLines[field] = struct{}{}
				lines = append(lines, line)
			}
		}
	} else {
		for scanner.Scan() {
			lineString := scanner.Text()
			lines = append(lines, strings.Fields(lineString))
		}
	}
	return lines, nil
}

func main() {
	keyProvided := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "k" {
			keyProvided = true
		}
	})

	file := flag.Arg(0)

	var k int
	var err error

	outFile, err := os.Create("out.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
	defer outFile.Close()
	if keyProvided {
		k, err = strconv.Atoi(key)
		if err != nil || k < 1 {
			fmt.Fprintf(os.Stderr,
				"invalid number at field start: invalid count at start of '%s'\n", key)
			return
		}
		k--
	}
	if k == 0 && !numeric {
		lines, err := readStrings(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
		var strSlice sort.StringSlice = lines
		if reverse {
			sort.Sort(sort.Reverse(strSlice))
		} else {
			sort.Strings(lines)
		}
		for _, line := range lines {
			outFile.WriteString(line + "\n")
		}
		return
	}

	lines, err := readStringSlices(file, k)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	var data sort.Interface
	if numeric {
		data = numSort{lines, k}
	} else {
		data = stringSort{lines, k}
	}
	if reverse {
		data = sort.Reverse(data)
	}
	sort.Sort(data)
	for _, line := range lines {
		outFile.WriteString(strings.Join(line, " ") + "\n")
	}
}
