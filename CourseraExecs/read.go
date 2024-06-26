package CourseraExecs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains
the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have
a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

type Name struct {
	Fname string
	Lname string
}

func Read() {
	var fileName string
	fmt.Print("Enter the name of the text file with .txt: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	var names []Name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 2 {
			names = append(names, Name{Fname: parts[0], Lname: parts[1]})
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", name.Fname, name.Lname)
	}
	file.Close()
}
