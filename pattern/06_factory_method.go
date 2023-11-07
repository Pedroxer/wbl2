package main

import (
	"fmt"
)

func CreateManipulator(name string) Manipulator {
	switch name {
	case "Line":
		return newLineManipulator()
	case "Text":
		return newTextManipulator()
	}
	return nil
}

// интерфейс над concrete products
type Manipulator interface {
	DownClick()
	Drag()
	UpClick()
}

type LineManipulator struct {
}

func newLineManipulator() *LineManipulator {
	return &LineManipulator{}
}
func (lm *LineManipulator) DownClick() {
	fmt.Println("DownClick from lm")
}
func (lm *LineManipulator) Drag() {
	fmt.Println("Drag from lm")
}
func (lm *LineManipulator) UpClick() {
	fmt.Println("UpClick from lm")
}

type TextManipulator struct {
}

func newTextManipulator() *TextManipulator {
	return &TextManipulator{}
}
func (tm *TextManipulator) DownClick() {
	fmt.Println("DownClick from tm")
}
func (tm *TextManipulator) Drag() {
	fmt.Println("Drag from tm")
}
func (tm *TextManipulator) UpClick() {
	fmt.Println("UpClick from tm")
}

func main() {
	linem := CreateManipulator("Line")
	txtm := CreateManipulator("Text")
	linem.DownClick()
	linem.UpClick()
	txtm.DownClick()
	txtm.UpClick()
}
