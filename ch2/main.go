package main

import "fmt"

// Constates que se pueden acceder desde min
const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

// const z = 20 * 10

func main() {
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

	// Son equivalentes
	// var x  = 1
	// x  := 1

	// x = x + 1
	// y = "bye"
	fmt.Println(nameKey)

	fmt.Println(x)
	fmt.Println(y);
	x := 15;
	fmt.Println(x)
	x = 50
	fmt.Println(x)
	// fstring := 'f'


	// f := 1
	// changeValue(f)

}

// Una funcion no puede cambiar el valor de un parámetro
// func changeValue(x int){
// 	x = 1
// 	fmt.Println(x)
// }