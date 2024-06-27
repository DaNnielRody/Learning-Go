package main

import (
	"fmt"
)

/*
Arrays em Go são inicializados como 0. Se definir um array de 5 (var x [5]int), serão 5 zeros.
*/

func arrays() {

	// Temos alguns jeitos de declarar arrays. O jeito explícito:

	var y [5]int = [5]int{1, 2, 3, 4, 5} // var y = [5]int{1,2,3,4,5}

	// E o jeito implícito. [...] indica que ele espera receber os valores que você adicionar dentro dos colchetes.
	x := [...]int{1, 2, 3}

	// Neste foor lop, podemos usar diretamente o length do array como range para iterar
	for i, v := range x {
		fmt.Printf("index %d, value %d\n", i, v)
	}

	fmt.Println(x, y)
}
