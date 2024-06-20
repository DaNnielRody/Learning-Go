package main

import "fmt"

func printing() {

	name := "Joe"
	age := 18
	fmt.Printf("Hi, %s. Are you %d?\n", name, age)
}

func converting() {

	/*
		Nem sempre conseguimos converter o valor que queremos, por exemplo:
		var x int32 = 1
		var y int16 = 2
		x = y
		Isso com certeza falhará, porque o compilador os assimilam de maneiras diferentes.

		Para resolver isso, podemos usar o cast convertion:
	*/

	var x int32 = 1
	var y int16 = 2
	x = int32(y)

	// Os dados são os mesmos, mas os endereços de memória são diferentes
	fmt.Println(&y, &x)
	fmt.Println(y, x)

	// Se quisermos apontar para o mesmo lugar e ter o mesmo valor, usamos ponteiros diretamente

	var h int = 50
	var j *int = &h

	fmt.Println(h, *j)
	fmt.Println(&h, j)

	/*
		Float:
		float32 -> 6 digitos de precisão
		float64 -> 15 digitos de precisão
	*/

}
