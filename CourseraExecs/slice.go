package CourseraExecs

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sliceMain() {
	slice := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin) // Essa entrada é como o Scanner scan = New Scanner(System.in) do Java

	for {
		fmt.Println("Enter an integer or 'X' to quit: ")

		/*
			Esses dois valores, correspondem ao scan.NextLine() do Java.
			Enquando no Java esse comando já lê a entrada e a retorna com esse valor, o Go precisa primeiro ler a linha, e depois retorná-la.
			Por isso scanner.Scan() -> para ler a linha
			E scanner.Text() -> Para armazenar o valor da linha e retornar como string
		*/
		scanner.Scan()
		input := scanner.Text()

		if strings.ToUpper(input) == "X" {
			break
		}

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter an integer or 'X' to quit.")
			continue
		}

		slice = append(slice, value)
		sort.Ints(slice)
		fmt.Println("Sorted slice:", slice)
	}
}
