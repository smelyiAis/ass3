package main

import (
    "fmt"
)

// Command interface defines the common methods for all concrete commands.
type Command interface {
    Execute()
}

// Receiver represents the object that the commands act upon.
type Light struct{}

func (l *Light) TurnOn() {
    fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
    fmt.Println("Light is off")
}

// ConcreteCommand for turning the light on.
type TurnOnCommand struct {
    light *Light
}

func (c *TurnOnCommand) Execute() {
    c.light.TurnOn()
}

// ConcreteCommand for turning the light off.
type TurnOffCommand struct {
    light *Light
}

func (c *TurnOffCommand) Execute() {
    c.light.TurnOff()
}

// Invoker stores and executes the command.
type RemoteControl struct {
    command Command
}

func (r *RemoteControl) SetCommand(command Command) {
    r.command = command
}

func (r *RemoteControl) PressButton() {
    r.command.Execute()
}

func main() {
    light := &Light{}
    turnOnCommand := &TurnOnCommand{light}
    turnOffCommand := &TurnOffCommand{light}

    remote := &RemoteControl{}

    remote.SetCommand(turnOnCommand)
    remote.PressButton()

    remote.SetCommand(turnOffCommand)
    remote.PressButton()
}
