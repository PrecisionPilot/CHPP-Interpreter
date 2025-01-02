package main

func main() {
	runeArray := []rune("可以asdf")
	// print(string(runeArray))

	for i := range len(runeArray) {
		println(string(runeArray[i]))
	}
}
