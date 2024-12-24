package main

type Car struct {
	name string
	mileage int
}

func (c Car) printName() {
	print(c.name)
}

func (c Car) honk() {
	print("pdiuaerui")
}



func main() {
	s := "a对"
	// runeArray := []rune(s)[5]
	print(len(s))


	familyCar *int
	somedudesCar Car
	familyCar.printName()
}
