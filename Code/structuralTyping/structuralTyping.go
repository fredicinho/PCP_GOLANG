package structuralTyping

type Meter struct {
	value int64
}

type Centimeter struct {
	value int64
}
func main() {
	cm := Centimeter{
		value: 1000,
	}

	var m Meter
	//allowed, because int32 and int 32 are
	m = Meter(cm)
	print(m.value)
	//same here
	cm = Centimeter(m)
	print(cm.value)
}
