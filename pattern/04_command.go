package main

import "fmt"

// invoker
type button struct {
	command Command
}

func NewButton(command Command) *button {
	return &button{
		command: command,
	}
}

func (b *button) Press() {
	b.command.Execute()
}

// command
type Command interface {
	Execute()
}

// concrete command 1
type onCommand struct {
	device Device
}

func NewOnCommand(device Device) *onCommand {
	return &onCommand{
		device: device,
	}
}

func (c *onCommand) Execute() {
	c.device.On()
}

// concrete command 2
type offCommand struct {
	device Device
}

func NewOffCommand(device Device) *offCommand {
	return &offCommand{
		device: device,
	}
}

func (c *offCommand) Execute() {
	c.device.Off()
}

// receiver
type Device interface {
	On()
	Off()
}

// concrete receiver
type tv struct {
	isRunning bool
}

func NewTv() *tv {
	return &tv{}
}

func (t *tv) On() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) Off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// client
func main() {
	tv1 := NewTv()
	onCommand := NewOnCommand(tv1)
	offCommand := NewOffCommand(tv1)
	onButton := NewButton(onCommand)
	onButton.Press()
	offButton := NewButton(offCommand)
	offButton.Press()

}
