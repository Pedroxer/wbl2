package main

import "testing"

func TestUnpackEmptyString(t *testing.T) {
	res, err := unpack("")
	if res != "" || err != nil {
		t.Fatalf(`unpack("") = %q, %v, expected "", nil`, res, err)
	}
}

func TestUnpackCorrectStringNoEscape(t *testing.T) {
	s := "a4bc2d5e"
	expected := "aaaabccddddde"
	res, err := unpack(s)
	if res != expected || err != nil {
		t.Fatalf(`unpack("") = %q, %v, expected %q, nil`, res, err, expected)
	}
}

func TestUnpackCorrectStringEscape(t *testing.T) {
	s := `qwe\4\5\\3\33`
	expected := `qwe45\\\333`

	res, err := unpack(s)
	if res != expected || err != nil {
		t.Fatalf(`unpack("") = %q, %v, expected %q, nil`, res, err, expected)
	}
}

// incorrect inputs
func TestUnpackStringStartWithANumber(t *testing.T) {
	res, err := unpack("45")

	if res != "" || err == nil {
		t.Fatalf(`unpack("45") = %q, %v, expected "", error`, res, err)
	}

}

func TestUnpackStringWithZero(t *testing.T) {
	res, err := unpack("a0b")

	if res != "" || err == nil {
		t.Fatalf(`unpack("45") = %q, %v, expected "", error`, res, err)
	}
}

func TestUnpackWrongEscape(t *testing.T) {
	res, err := unpack(`qwe\\\`)

	if res != "" || err == nil {
		t.Fatalf(`unpack("45") = %q, %v, expected "", error`, res, err)
	}
}
