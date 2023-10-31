package commands

import (
	"main/command/receiver"
)

type LightOffCommand struct {
	light *receiver.Light
}

func NewLightOffCommand(light *receiver.Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (loc *LightOffCommand) Execute() {
	loc.light.TurnOff()
}
