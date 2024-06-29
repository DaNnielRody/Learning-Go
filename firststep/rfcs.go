package main

import (
	"encoding/json"
	"fmt"
)

func jsons() {

	type Person struct {
		Name  string
		Addr  string
		Phone string
	}

	p1 := Person{Name: "Dan", Addr: "Rua dois", Phone: "2436235425"}

	barr, err := json.Marshal(p1) // Transforma em arquivo json
	if err != nil {
		fmt.Println("We can't convert into json")
		return
	}

	fmt.Println(string(barr)) // Lê o arquivo json como string

	var p2 Person

	err = json.Unmarshal(barr, &p2) // Pega o valor em Json de p1 e copia em p2

	fmt.Println(p2)
}

func files() {

}
