package main

import (
	"reflect"
	"testing"
)

func TestStringToFloat(t *testing.T) {
	f, ok := stringToFloat("3.14")
	if f != 3.14 || !ok {
		t.Logf(`stringToFloat("3.14") = %f, %v, expected 3.14, true`, f, ok)
		t.Fail()
	}

	f, ok = stringToFloat("qwerty")
	if f != 0 || ok {
		t.Logf(`stringToFloat("qwerty") = %f, %v, expected 0, false`, f, ok)
		t.Fail()
	}

	f, ok = stringToFloat("nan")
	if f != 0 || ok {
		t.Logf(`stringToFloat("nan") = %f, %v, expected 0, false`, f, ok)
		t.Fail()
	}
}

func TestLessNum(t *testing.T) {
	less := lessNum("0.1", "1.0")
	if !less {
		t.Logf(`stringToFloat("0.1", "1.0") = %v, expected true`, less)
		t.Fail()
	}

	less = lessNum("0.1", "-1.0")
	if less {
		t.Logf(`stringToFloat("0.1", "-1.0") = %v, expected false`, less)
		t.Fail()
	}

	less = lessNum("0.1", "0.1")
	if less {
		t.Logf(`stringToFloat("0.1", 0.1") = %v, expected false`, less)
		t.Fail()
	}

	less = lessNum("qwe", "1.0")
	if !less {
		t.Logf(`stringToFloat("qwe", "1.0") = %v, expected true`, less)
		t.Fail()
	}

	less = lessNum("qwe", "asd")
	if less {
		t.Logf(`stringToFloat("qwe", "asd") = %v, expected false`, less)
		t.Fail()
	}
}

func TestReadStringsUnique(t *testing.T) {
	unique = true
	file := "test.txt"

	expected := []string{
		"6 7 3",
		"January 8 dog",
		"11 8 0",
		"March 8 -1",
		"6 7 3 5 cat 12",
	}

	lines, err := readStrings(file)
	if err != nil {
		t.Fatalf("readStrings: %v", err)
	}
	if !reflect.DeepEqual(lines, expected) {
		t.Fatal("result of readStrings is differ from expected")
	}
}

func TestReadStringSlicesUniqueKey(t *testing.T) {
	unique = true
	k := 1
	file := "test.txt"

	expected := [][]string{
		{"6", "7", "3"},
		{"January", "8", "dog"},
	}

	lines, err := readStringSlices(file, k)
	if err != nil {
		t.Fatalf("readStringSlices: %v", err)
	}
	if !reflect.DeepEqual(lines, expected) {
		t.Fatal("result of readStringSlices is differ from expected")
	}
}
