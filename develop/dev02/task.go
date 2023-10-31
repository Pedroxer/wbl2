package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

func unpack(input string) (string, error) {
	if len(input) == 0 {
		return input, nil
	}

	ir := []rune(input)
	var res []rune
	if unicode.IsDigit(ir[0]) { // если строка начинается с числа - она некорректная
		return "", errors.New("unpack: string starts with digit")
	}
	for i := 0; i < len(ir); i++ {
		switch {
		case unicode.IsDigit(ir[i]):
			n := int(ir[i] - '0')

			if n == 0 {
				return "", errors.New("unpack: found 0 in string")
			}
			for i < len(ir)-1 && unicode.IsDigit(ir[i+1]) {
				n = n*10 + int(ir[i+1]-'0') // пока следующая руна является числом, мы продолжаем высчитывать n
				i++
			}
			pack := make([]rune, n-1)
			for i := range pack {
				pack[i] = res[len(res)-1]
			}
			res = append(res, pack...)
			if n == 1 {
				break
			}
		case ir[i] == '\\':

			// одиночный '\' в конце строки некорректен
			if i == len(ir)-1 {
				return "", errors.New("unpack: incorrect escape seq")
			}

			// руна после '\' обрабатывается как символ
			res = append(res, ir[i+1])
			i++
		default:
			res = append(res, ir[i])
		}
	}
	return string(res), nil
}

func main() {
	var input string
	fmt.Scan(&input)
	str, err := unpack(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(str)

}
