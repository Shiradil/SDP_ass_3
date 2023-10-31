package receiver

type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	println("Light is ON")
}

func (l *Light) TurnOff() {
	l.isOn = false
	println("Light is OFF")
}
