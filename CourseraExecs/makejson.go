package CourseraExecs

/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func MakeJson() {

	people := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter your address: ")
	scanner.Scan()
	address := scanner.Text()

	people["name"] = name
	people["address"] = address

	js, err := json.Marshal(people)
	if err != nil {
		fmt.Println("We can't convert into jason", err)
		return
	}

	fmt.Println(js)         // byte storage version
	fmt.Println(string(js)) // json format version
}
