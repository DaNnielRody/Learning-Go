package main

import "fmt"

/*
Hello World e Ponteiros
*/

func main() {
	fmt.Println("Hello World")

	count := 1
	p := &count
	// Para criarmos uma variável apontando diretamente para algum tipo, temos duas maneiras:
	var ip *int     // Aqui ele aponta para um inteiro qualquer, mas ainda não recebe um valor.
	ptr := new(int) // Aqui ele cria uma variável apontando para o endereço de memória dela. Sabemos que uma variável não inicializada é 0, então se criarmos assim...

	fmt.Println(ptr, *ptr) // Vemos o endereço de memória que referencia 0, e o zero sendo desreferenciado e exibido.

	/*
		Caso eu queira alocar um novo valor, isso tem que ser feito assim:
		Como o ptr é um endereço apontando para o ponteiro, tenho que alterar diretamente o seu dado. ptr = 3 não daria certo.
		E mesmo se eu quiser imprimir, tenho que usar o asterisco, caso contrário ele vai continuar apontando para o endereço.
	*/
	*ptr = 3
	fmt.Println(ptr, *ptr)

	// Já desta forma, quando iniciamos a variável com o ponteiro para alguma referência de inteiro, ele não recebe valor.
	fmt.Println(ip) // O resultado é nil, pois ip não aponta para nenhum endereço válido.

	ip = p             // Porém, se apontarmos para um espaço existente, ele recebe o valor do endereço em questão, que neste caso é p.
	fmt.Println(ip, p) // Mais uma vez. Se neste caso p não fosse um endereço, eu precisaria iniciá-lo como ip = &p

	// E se quisermos que ele exiba o valor do dado, é só utilizar o asterisco para desreferenciar o endereço de memória.
	fmt.Println(*ip, *p)

	// Resumindo:
	fmt.Println("O endereço é:", p)    // Vemos o endereço porque p = &count (endereço da variável count).
	fmt.Println("O dado de *p é:", *p) // Vemos o valor porque *p desreferencia p, dando o valor armazenado no endereço de count.

	// Prints

	printing()
	converting()
	pack()
	//forLoops()
	switches()
	//read()
	slices()
}
