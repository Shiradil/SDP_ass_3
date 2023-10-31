package commands

import (
	"main/command/receiver"
)

type LightOnCommand struct {
	light *receiver.Light
}

func NewLightOnCommand(light *receiver.Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (loc *LightOnCommand) Execute() {
	loc.light.TurnOn()
}
