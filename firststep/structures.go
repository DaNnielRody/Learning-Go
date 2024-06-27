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

/*
Slices são pequenas janelas, que se separam de uma vista maior (o array)
*/
func slices() {

	/*
		O slice determina um ponto para começar, e vai até um número anterior último valor.
		O slice pode adicionar valores, mas seu limite é o exclusivo(quando não especificado, é o próprio array).
		Por exemplo, se você cria um slice a partir do index 2 e o array tem 100 elementos, ele só poderá armazenar mais 98
	*/

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	// Slice sem exclusivo:
	slice1 := arr[1:3] // b, c
	slice2 := arr[2:5] // c, d, e

	fmt.Println(cap(slice1)) // 25 -> Capacidade total
	fmt.Println(cap(slice2)) // 24 -> Como começa no index 2, ele só pode adicionar a partir daí
	// Slice com exclusivo:
	slice3 := arr[1:4:6]     // Com esse último parâmetro, o slice permite que mais elementos sejam adicionados sem a necessidade de alocação(até o limite de 7).
	fmt.Println(cap(slice3)) // É 5 porque o exclusivo é 6 e o index é 1. Então ele faz a conta: 6 - 1 = 5

	// Agora vamos aumentar a capacidade:
	slice3 = append(slice3, slice2...)

	fmt.Println(cap(slice3)) // Como o slice passou do limite, o Go dobra o tamanho que ele suporta para guardar mais valores. Como antes era 5, agora são 10.

	slice3 = append(slice3, slice1...)
	slice3 = append(slice3, slice1...)
	slice3 = append(slice3, slice1...) // Como aqui ele passa novamente, mais uma vez ele dobra, passando de 10 para 20
	fmt.Println(cap(slice3))

	/*
		Aqui, podemos ver o slice3 recebendo parâmetros além do limite. Por que isso ocorre e qual a diferença?
		Quando não definimos o limite (como o exemplo dos primeiros slices), o Go usa a capacidade até o fim do array.
		Porém, ao delimitar, você especifica até onde ele pode crescer sem precisar realocar. Caso passe disso, ele realoca automaticamente.
		Delimitando o slice, você consegue poupar memória e tornar a gestão do código mais eficiente.
	*/

	fmt.Println(slice1, slice2, slice3)

	sli := []int{1, 2, 3} // Aqui se cria um slice literal
	fmt.Println(sli)
}
