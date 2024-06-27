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
		O slice é a forma dinâmica de adicionar valores num array. Com o array, adicionamos apenas tamanho, enquanto no slice adicionamos os valores.
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
	// Caso não tivéssemos valores definidos dentro do array, apenas o tamanho, poderíamos adicionar assim: slice3 = append(slice3, "b", "d", "k", ...)

	fmt.Println(cap(slice3)) // Como o slice passou do limite, o Go dobra o tamanho que ele suporta para guardar mais valores. Como antes era de 5, agora é de 10.

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

	sli := []int{1, 2, 3}       // Aqui se cria um slice literal
	sli2 := make([]int, 10, 15) // Aqui se cria um slice com o make, delimitando o tamanho de 10, mas que tem a capacidade até 15
	fmt.Println(sli, sli2)
}

func mapT() {
	/*
		No map, assim como no Hash, utilizamos a key para retornar o valor, recebemos: [key]value.
		Quando iniciamos o map vazio, precisamos utilizar a sintaxe com o make:
	*/
	var idMap map[string]int
	idMap = make(map[string]int) // idMap := make(map[string]int)

	// Depois, criamos os chave-valores:

	idMap["John"] = 13
	idMap["Alice"] = 14

	/*
		No map, também temos como "buscar" um valor. Passamos os parâmetros value e boolean.
		O primeiro vaor é associado ao valor do map, e o segundo verifica se ele está presente no map ou não.
		No caso, é como se estivéssemos fazendo uma busca. Se tentarmos imprimir o id, é a mesma coisa que pesquisar idMap["Alice"]:
	*/

	fmt.Println(idMap["Alice"]) // 14

	// Fazemos referência:
	id, p := idMap["Alice"]

	// E buscamos:
	fmt.Println(id, p) // 14 true

	// Mas e se deletá-lo?
	delete(idMap, "Alice")

	/*
		Nessa referência, quando deletamos, precisamos fazer uma nova busca para verificar se o valor está dentro do map.
		Isto ocorre pois os valores são determinados a partir do momento em que você faz a busca (diferentemente de usar idMap["Alice"] diretamente).
		Se tentarmos printar antes de fazer uma nova busca, ele ainda exibe os valores anteriores:
	*/

	fmt.Println(idMap["Alice"], id, p) // 0 14 true

	// Atualizamos a busca:
	id, p = idMap["Alice"]

	// Agora, se buscarmos, ela atualizou:
	fmt.Println(idMap["Alice"], id, p) // 0 0 false

	/*
		Podemos inicializar o map com valores utilizando a sua expressão literal.
		Como nela já temos valores, não precisamos utilizar o "make"
	*/

	idMap2 := map[string]int{
		"John": 12,
	}

	fmt.Println(idMap2)

	// Podemos adicionar os valores de volta:

	idMap["Anthony"] = 15
	idMap["Eduarda"] = 19
	idMap["Daniel"] = 21

	// Por fim, podemos iterar para exibir os valores:

	for k, v := range idMap {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}

func structs() {

	// Em Go não temos objeto, temos estruturas. Elas são declaradas assim:
	type Person struct {
		name    string
		address string
		phone   int
	}

	// A forma normal:
	// Desse jeito, qualquer mudança dentro do struct não vai alterar o struct original, porque você trabalha apenas com cópias do struct.
	var person2 Person
	person2.phone = 23
	person2.name = "Oi"
	person2.address = "wriktekr"

	// A com ponteiros:
	// Desta forma, vocẽ altera diretamente o ponteiro original, o que pode ser útil para gerenciar memória.
	personPtr := new(Person)
	personPtr.name = "Pt"
	personPtr.address = "Ponteiro"
	personPtr.phone = 9047140

	// E temos a forma literal:
	// Esta é uma forma conveniente de inicializar um struct com valores específicos.
	person := Person{name: "Dan", address: "Rua do zap", phone: 40038922}
	fmt.Println(person, person2, *personPtr)

	/*
		Resumindo:
		Use Inicialização Normal quando quiser inicializar um struct com valores padrão ou desconhecidos no momento da declaração.
		Use Inicialização com Ponteiros quando precisar passar o struct por referência para evitar cópias ou quando precisar gerenciar manualmente a alocação de memória.
		Use Inicialização Literal quando souber os valores dos campos antecipadamente e quiser uma forma concisa e direta de inicializar o struct.
	*/
}
