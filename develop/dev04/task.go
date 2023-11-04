package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortAnagram(input *[]string) *map[string]*[]string {
	anagrmas := make(map[string][]string)
	res := make(map[string]*[]string)
	for _, val := range *input {
		v := strings.ToLower(val)
		sortedV := sortString(v)
		anagrmas[sortedV] = append(anagrmas[sortedV], v)
	}

	for _, v := range anagrmas {
		v := v
		if len(v) > 1 {
			res[v[0]] = &v
			sort.Slice(v, func(i, j int) bool {
				return v[i] < v[j]
			})
		}
	}
	return &res
}

func sortString(s string) string {
	runeString := []rune(s)
	sort.Slice(runeString, func(i, j int) bool {
		return runeString[i] < runeString[j]
	})
	return string(runeString)
}

func main() {
	s := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Кролик"}
	res := sortAnagram(&s)
	for key, value := range *res {
		fmt.Println(key, *value)
	}
}
