package main

import (
	"reflect"
	"regexp"
	"testing"
)

func TestReadStringsFromFile(t *testing.T) {
	file := "file.txt"
	expected := []string{
		"0123",
		"0123",
		"abc",
		"123123a",
		"[",
		"[]",
		"абсд",
	}

	lines, err := readStringsFromFile(file)
	if err != nil {
		t.Fatalf("readStringsFromFile: %v", err)
	}
	if !reflect.DeepEqual(lines, expected) {
		t.Fatal("result of readStrings is differ from expected")
	}
}

func TestPrintWithAfterWithoutLineNumber(t *testing.T) {
	lines := []string{
		"0123",
		"abc",
		"123123a",
		"[",
		"[]",
		"абсд",
	}
	expected := []string{
		"abc",
		"123123a",
		"[",
		"--",
	}
	after = "3"
	reg, err := regexp.Compile("0123")
	if err != nil {
		t.Fatalf("printWithAfter: %v", err)
	}
	str := printWhenAfter(reg, lines, 0)
	if !reflect.DeepEqual(str, expected) {

		t.Fatal("result of readStrings is differ from expected")
	}
}

func TestPrintWithAfterWtihLineNumber(t *testing.T) {
	lines := []string{
		"0123",
		"abc",
		"123123a",
		"[",
		"[]",
		"абсд",
	}
	expected := []string{
		"Line 2: abc",
		"Line 3: 123123a",
		"Line 4: [",
		"--",
	}
	after = "3"
	lineNum = true
	reg, err := regexp.Compile("0123")
	if err != nil {
		t.Fatalf("printWithAfter: %v", err)
	}
	str := printWhenAfter(reg, lines, 0)
	if !reflect.DeepEqual(str, expected) {

		t.Fatal("result of readStrings is differ from expected", str)
	}
}
