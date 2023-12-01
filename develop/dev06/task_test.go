package main

import (
	"math"
	"reflect"
	"testing"
)

func TestParseFiledsCorrectInput(t *testing.T) {
	s := "-2,5,7-9,12-"
	expected := [][]int{{1, 2}, {5, 5}, {7, 9}, {12, math.MaxInt}}

	parsed, err := parseFields(s)

	if err != nil || !reflect.DeepEqual(parsed, expected) {
		t.Logf("parseFields(%s) = %v, %v, expected: %v, nil", s, parsed, err, expected)
	}
}

func TestParseFiledsIncorrectInputDecreasingRange(t *testing.T) {
	s := "5-2"

	parsed, err := parseFields(s)

	if err == nil {
		t.Logf("parseFields(%s) = %v, nil, expected: nil, error", s, parsed)
	}
}

func TestParseFiledsIncorrectInputFormat(t *testing.T) {
	s := "hello"

	parsed, err := parseFields(s)

	if err == nil {
		t.Logf("parseFields(%s) = %v, nil, expected: nil, error", s, parsed)
	}
}

func TestParseFiledsIncorrectInputComma(t *testing.T) {
	s := ","

	parsed, err := parseFields(s)

	if err == nil {
		t.Logf("parseFields(%s) = %v, nil, expected: nil, error", s, parsed)
	}
}

func TestParseFiledsIncorrectInputPreComma(t *testing.T) {
	s := ",2"

	parsed, err := parseFields(s)

	if err == nil {
		t.Logf("parseFields(%s) = %v, nil, expected: nil, error", s, parsed)
	}
}

func TestParseFiledsIncorrectInputPostComma(t *testing.T) {
	s := "2,"

	parsed, err := parseFields(s)

	if err == nil {
		t.Logf("parseFields(%s) = %v, nil, expected: nil, error", s, parsed)
	}
}

func TestIndexInSegmentsTrue(t *testing.T) {
	i := 8
	segs := [][]int{{7, 9}}

	res := indexInSegments(i, segs)

	if !res {
		t.Logf("indexInSegments(%d, %v) = %v, expected: true", i, segs, res)
	}
}

func TestIndexInSegmentsFalse(t *testing.T) {
	i := 5
	segs := [][]int{{7, 9}}

	res := indexInSegments(i, segs)

	if res {
		t.Logf("indexInSegments(%d, %v) = %v, expected: true", i, segs, res)
	}
}
