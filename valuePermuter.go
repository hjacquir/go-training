package main

import "fmt"

type StringManipulator struct {
	A string
	B string
}

// sans le * la permutation ne fonctionnerait pas -> on aurait {A B} au lieu de {B A}
func (manipulator *StringManipulator) permut() {
	var saveA string = manipulator.A
	var saveB string = manipulator.B

	manipulator.A = saveB
	manipulator.B = saveA
}

func main()  {
	var word StringManipulator = StringManipulator{"A", "B"}
	word.permut()
	fmt.Println(word)
}
