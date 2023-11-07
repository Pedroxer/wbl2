package main

import "fmt"

// builder
type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

// product
type House struct {
	windowType string
	doorType   string
	floor      int
}

func NewHouse(windowType string, doorType string, floor int) House {
	return House{
		windowType: windowType,
		doorType:   doorType,
		floor:      floor,
	}
}

func (h *House) GetWindowType() string {
	return h.windowType
}

func (h *House) GetDoorType() string {
	return h.doorType
}

func (h *House) GetFloor() int {
	return h.floor
}

// concrete builder for igloo
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) SetWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) SetNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) GetHouse() House {
	return NewHouse(b.windowType, b.doorType, b.floor)
}

// concrete builder for normal house
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) SetWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) SetNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) GetHouse() House {
	return NewHouse(b.windowType, b.doorType, b.floor)
}

// director
type director struct {
	builder Builder
}

func NewDirector(b Builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *director) BuildHouse() House {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}

func getBuilder(builderType string) Builder {
	if builderType == "normal" {
		return NewNormalBuilder()
	}
	if builderType == "igloo" {
		return NewIglooBuilder()
	}
	return nil
}

func main() {
	nBuilder := getBuilder("normal")
	igBuilder := getBuilder("igloo")

	direct := NewDirector(nBuilder)
	normalHouse := direct.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.GetFloor())

	direct.SetBuilder(igBuilder)
	iglooHouse := direct.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.GetFloor())
}
