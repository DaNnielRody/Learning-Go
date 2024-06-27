package CourseraExecs

import (
	"fmt"
	"strings"
)

/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or
lower-case.
*/

func checkString(s string) bool {
	strings.ToLower(s)
	return strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a")
}

func main() {

	var enteringString string

	fmt.Println("Type a string: ")
	_, err := fmt.Scan(&enteringString)

	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	if checkString(enteringString) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found!")
	}
}
