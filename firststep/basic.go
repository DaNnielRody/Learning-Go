package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Como no unicode é incluso o ASCII, onde cada conjunto de bits representa um caracter, ele precisa receber esses caracteres
para fazer a verificação. No go, cada caracter unicode é representado por runa, permitindo que o go represente qualquer
ponto de código Unicode. Por isso, para fazer a verificação, precisamos utilizar a seguinte sintaxe:
*/

func pack() {
	i := '9'
	j := ' '
	w := "a"
	z := "zap zop zup"
	x := "12"
	y := 1
	fmt.Println(unicode.IsDigit(i), unicode.IsSpace(j))
	fmt.Println(strings.ToUpper(w))
	fmt.Println(strings.Contains(z, "zep"))
	fmt.Println(strings.TrimSpace(z))
	// Aqui retorna a conversão e o erro:
	fmt.Println(strconv.Atoi(x)) // Vira int
	fmt.Println(strconv.Itoa(y)) // Vira string

}

func constans() {
	/*
		Temos dois jeitos de declarar constantes.
		O jeito tradiciona:
	*/

	const x int = 25

	// É o jeito "enum" (em outras linguagens de programação)

	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		E
		F
	)

	/*
		Do jeito "iota", eu declaro apenas que quero que os valores sejam diferentes, não importando quais sejam seus valores.
		Por isso é parecido com os tipos Enum, por exemplo.
		Utilizamos esses tipos para meses, notas, dias da semana, etc...
	*/
}

func forLoops() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/*
		Este modo é como se fosse um while:
	*/

	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
}

/*
O switch é como um if statement, que retorna um booleano. Ele pega a tag que é para ser checada e compara com cada valor de cada caso.
Se a tag checaca é encontrada, a expressão é executada
*/

func switches() {

	const text = "hello world"

	switch text {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	case "hello world":
		fmt.Println(text)
	default:
		fmt.Println("not found")
	}

	/*
		Também temos o switch sem tag, que pode ser utilizado exatamente como um if-else statement.
		O uso do switch é recomendado em casos onde as condições são maiores e mais relacionadas (como por exemplo, exibição de notas).
		O uso do if-else é mais recomendado para casos diretos, ou onde as condições não estão muito relacionadas.
	*/
	x := 0

	/*
		Aqui, o primeiro caso com a condição como verdadeira, retornará
	*/
	switch {
	case x > 1:
		fmt.Println("x > 1")
	case x < 1:
		fmt.Println("x < 1")
	default:
		fmt.Println("other")

	}
}

func read() {

	/*
		A função Scan lê a entrada do usuário e armazena o valor nas variáveis fornecidas.
		Para isso, é necessário passar o endereço de memória (ponteiro) dessas variáveis,
		permitindo que Scan possa modificar o valor diretamente no local onde ele está armazenado.
	*/

	var apple int

	fmt.Println("Number of Apples?")
	_, err := fmt.Scan(&apple)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println(apple)

	/*
		É o mesmo que isso:
		apple := new(int) // Aponta diretamente para um endereço

			fmt.Println("Number of Apples?")
			_, err := fmt.Scan(apple) // Aponta o valor digitado pelo usuário para este endereço

			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			fmt.Println(*apple) // Exibe o valor
	*/
}
