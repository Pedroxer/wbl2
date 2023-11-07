package main

import (
	"fmt"
	"log"
)

// Context
type VendingMachine struct {
	HasItem       State
	ItemRequested State
	HasMoney      State
	NoItem        State

	currentState State

	ItemCount int
	ItemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		ItemCount: itemCount,
		ItemPrice: itemPrice,
	}

	hasItemState := NewHasItemState(v)
	itemRequestedState := NewItemRequestedState(v)
	hasMoneyState := NewHasMoneyState(v)
	noItemState := NewNoItemState(v)

	v.SetState(hasItemState)
	v.HasItem = hasItemState
	v.ItemRequested = itemRequestedState
	v.HasMoney = hasMoneyState
	v.NoItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.ItemCount = v.ItemCount + count
}

// State
type State interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}

// Concrete state 1
type noItemState struct {
	vendingMachine *VendingMachine
}

func NewNoItemState(vendingMachine *VendingMachine) *noItemState {
	return &noItemState{
		vendingMachine: vendingMachine,
	}
}

func (i *noItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) AddItem(count int) error {
	i.vendingMachine.IncrementItemCount(count)
	i.vendingMachine.SetState(i.vendingMachine.HasItem)
	return nil
}

func (i *noItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}

// Concrete state 2
type hasItemState struct {
	vendingMachine *VendingMachine
}

func NewHasItemState(vendingMachine *VendingMachine) *hasItemState {
	return &hasItemState{
		vendingMachine: vendingMachine,
	}
}

func (i *hasItemState) RequestItem() error {
	if i.vendingMachine.ItemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.NoItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requested\n")
	i.vendingMachine.SetState(i.vendingMachine.ItemRequested)
	return nil
}

func (i *hasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.IncrementItemCount(count)
	return nil
}

func (i *hasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *hasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}

// Concrete state 3
type itemRequestedState struct {
	vendingMachine *VendingMachine
}

func NewItemRequestedState(vendingMachine *VendingMachine) *itemRequestedState {
	return &itemRequestedState{
		vendingMachine: vendingMachine,
	}
}

func (i *itemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *itemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.ItemPrice {
		return fmt.Errorf("insert money is less. Please insert %d", i.vendingMachine.ItemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.SetState(i.vendingMachine.HasMoney)
	return nil
}

func (i *itemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}

// Concrete state 4
type hasMoneyState struct {
	vendingMachine *VendingMachine
}

func NewHasMoneyState(vendingMachine *VendingMachine) *hasMoneyState {
	return &hasMoneyState{
		vendingMachine: vendingMachine,
	}
}

func (i *hasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *hasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.ItemCount = i.vendingMachine.ItemCount - 1
	if i.vendingMachine.ItemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.NoItem)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.HasItem)
	}
	return nil
}
func main() {
	vendingMachine := NewVendingMachine(1, 10)
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
