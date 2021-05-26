package main

type Meter struct {
	value int32
}

type Centimeter struct {
	value int64
}

func main() {
	cm := Centimeter{
		value: 1000,
	}

	var m Meter

	m = Meter(cm)
	println(m.value)
	//same here
	cm = Centimeter(m)
	print(cm.value)
}
