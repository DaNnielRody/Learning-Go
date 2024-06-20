package main

import (
	"fmt"
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
	fmt.Println(unicode.IsDigit(i), unicode.IsSpace(j))
	fmt.Println(strings.ToUpper(w))
	fmt.Println(strings.Contains(z, "zop"))
}
