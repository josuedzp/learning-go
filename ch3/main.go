package main

func main() {

	// Arrays

	// Declarando un array sin valores;
	// var x [3]int // results: [0,0,0]

	// Asignar valor a array inicializado en 0;
	// x[0] = 1

	// Declarando un array longitud fija con valores;
	// var x = [3]int {10,20,30}

	// Declarando un array de longitug 12 e indicar que posiciones tendrán valor;
	// var x = [12]int{1,2,3} // posiciones 0 -> 1, 1 -> 2, 2 -> 3, las demás seran inicializadas a 0
	// var x = [12]int{1, 5:4, 6, 10:100} // posiciones 0 -> 1, 5 -> 4, 6 -> 6, 10 -> 100, las demás seran inicializadas a 0

	// Declarando un array de longitud fija una hecha la asignación
	// var x = [...]int{10,20,30,40} 

	// Comparación de arrays
	// var x = [...]int{1, 2, 3}
	// var y = [3]int{1, 2, 3}
	// fmt.Println(x == y) // prints true

	// Instancia matriz
	// var x [2][3]int

	// Al intentar acceder a un indice no exitente falla en tiempo de compilación
	// x := [...]int {10,20}
	// fmt.Println(x[2])


	// Longitud de array
	// var x[2]int
	// len(x)

	/* 
		Array tiene longitud fija
		Slices tienen longitud variable
	*/

	// Array [...]
	// Slice []

	// Instancia un slice
	// var x = []int{10, 20, 30}

	// Especificar las posiciones de un slice
	// var x = []int{1, 5: 4, 6, 10: 100, 15}

	// Matriz de slice
	// var x [][] int

	// Un slice también falla en tiempo de compilación al intenctar acceder a un indice negativo o fuera de rango
	// var x = []int {1,2}
	// fmt.Println(x[2])

	// Diferencias entre slice y array
	// var x []int // results: []
	// Nil es un identificador que representa la valte de valor de algun tipo, es ligeramente diferente de un nulo en otros idiomas
	// fmt.Println(x == nil) // prints true

	
	// Append: añadir elementos al slice
	// var x []int
	// x = append(x, 10) // añade 1 elemento
	// x = append(x, 10, 20, 30) // añade varios elementos

	// Desestructuración
	// var x []int
	// y := []int{20, 30, 40}
	// x = append(x, y...)

	// Capacidad de longitud de un slice, el slice a medida que se va añadiendo y llega a la capacidad el próximo append duplicará la capacidad
	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// fmt.Println(x, len(x), cap(x))

	// Declarando slice vacío con longitud definida
	// x := make([]int, 5)
	// x = append(x, 10)
	// x = append(x, 10) // el resultado no será igual a longitud 7 y capacidad 7

	// Declarando slice vacío con longitud y capacidad definida
	// x := make([]int, 5, 7)
	// x = append(x, 10)
	// x = append(x, 10) // el resultado será igual a longitud 7 y capacidad 7

	// fmt.Printf("los valores son: %v, su longitud es de: %v, su capacidad es de:% v\n",x, len(x), cap(x))

	// Declarando un slice que podría quedarse nil
	// var data []int
	// if data == nil {
	// 	fmt.Println("data is nil") // prints data is nil
	// }

	// Declarando un slice que no es nil
	// var data2 = []int{}
	// if data2 == nil {
	// 	fmt.Println("data is nil")// no prints 

	// Declarando un slice con valores por defecto
	// data := []int{2,4,6,8}


	// Slicing slices
	// x := []int{1, 2, 3, 4}
	// y := x[:2]
	// z := x[1:]
	// d := x[1:3]
	// e := x[:]
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)
	// fmt.Println("d:", d)
	// fmt.Println("e:", e)

	// Un slice de un slice no hace una copia, los slices están compartiendo la misma memoria
	// x := []int{1, 2, 3, 4}
	// y := x[:2]
	// z := x[1:]
	// x[1] = 20
	// y[0] = 10
	// z[1] = 30
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// Un slice de un slice también comparte la capacidad del original
	// x := []int{1, 2, 3, 4}
	// y := x[:2]
	// fmt.Println(cap(x), cap(y))
	// y = append(y, 30) // Esto añadirá un elemento a la ultima posición de 'y' pero también la añadirá/remplezará a la posición que corresponda del slice original
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	// Un slice de un slice puede llegar a complicarse mucho
	// x := make([]int, 0, 5)
	// x = append(x, 1, 2, 3, 4) // [1,2,3,4]
	// fmt.Println(cap(x)) 
	// y := x[:2] // [1,2]
	// fmt.Println(cap(x), cap(y)) 
	// z := x[2:] // [3,4] y tomará la capacidad restante
	// fmt.Println(cap(x), cap(y), cap(z)) 
	// y = append(y, 30, 40, 50) // [1,2,30,40,50]
	// x = append(x, 60) // [1,2,30,40,60]
	// z = append(z, 70) // [30,40,70]
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// Slice de un slice con capacidad
	// x := make([]int, 0, 5)
	// x = append(x, 1, 2, 3, 4)
	// y := x[:2:2] // se puede hacer un slice con una capacidad reducida
	// z := x[2:4:4] // no se puede hacer un slice con una capacidad mayor
	// fmt.Println(cap(x), cap(y), cap(z)) 
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)


	// fmt.Printf("los valores son: %v, su longitud es de: %v, su capacidad es de:% v\n",data, len(data), cap(data))

}