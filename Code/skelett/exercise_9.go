package main

import "fmt"

type human struct {
	age        float32
	gender     string
	lengthBone float32
}

func blength(human human) float32 {
	if human.gender == "male" {
		return human.lengthBone*2.318 + 69.089 - ageValue(human)
	} else {
		return human.lengthBone*2.317 + 61.412 - ageValue(human)
	}
}

func ageValue(human2 human) float32 {
	if human2.age > 30 {
		return human2.age * 0.006
	} else {
		return 0
	}
}
func main() {
	human := human{age: 92, gender: "male", lengthBone: 50}
	fmt.Println(blength(human))
}
