package invoker

type Command interface {
	Execute()
}

type SimpleRemoteControl struct {
	slot Command
}

func (src *SimpleRemoteControl) SetCommand(command Command) {
	src.slot = command
}

func (src *SimpleRemoteControl) PressButton() {
	src.slot.Execute()
}
