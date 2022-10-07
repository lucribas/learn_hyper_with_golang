package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Legs int
	Age  int
	Name string
}

type Dog struct {
	Animal
	Latido string
}

func (d Dog) fazBarulho() {
	fmt.Println("barulho: ", d.Latido)
}

type Cat struct {
	Animal
	Ruido string
}

func (c Cat) fazBarulho() {
	fmt.Println("barulho: ", c.Ruido)
}

type Barulhento interface {
	fazBarulho()
}

func Som(a Barulhento) {
	a.fazBarulho()
}

func main() {

	// cat := Animal{Name: "fifi", Age: 4, Legs: 4}
	// dog := Animal{Name: "bob", Age: 10, Legs: 4}
	// cat_str, _ := json.Marshal(cat)
	// fmt.Println(string(cat_str))
	// dog_str, _ := json.Marshal(dog)
	// fmt.Println(string(dog_str))

	dog2 := Dog{Animal: Animal{Name: "tod", Age: 8, Legs: 4}, Latido: "auaua"}
	cat2 := Cat{Animal: Animal{Name: "mimi", Age: 2, Legs: 4}, Ruido: "meow"}

	dog2_str, _ := json.Marshal(dog2)
	fmt.Println(string(dog2_str))
	dog2.fazBarulho()

	cat2_str, _ := json.Marshal(cat2)
	fmt.Println(string(cat2_str))
	cat2.fazBarulho()

	Som(cat2)
	Som(dog2)
}
