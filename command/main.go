package main

import (
	"main/command/cmd/commands"
	"main/command/invoker"
	"main/command/receiver"
)

func main() {
	light := &receiver.Light{}

	lightOn := commands.NewLightOnCommand(light)
	lightOff := commands.NewLightOffCommand(light)

	remote := &invoker.SimpleRemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()
}
