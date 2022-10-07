// Examplo structs
package main

import (
	"fmt"
)

type Animal struct {
	Age int
	Leg int
}

func (a *Animal) Move() {
	fmt.Println("Um animal se move com ", a.Leg, " pernas.")
}

type Cachorro struct {
	Animal
	Leg int
}

func (d *Cachorro) Move() {
	fmt.Println("Um cachorro se cove com ", d.Leg, " pernas.")
}
func main() {
	animal := Animal{
		Age: 5,
		Leg: 2,
	}
	// Output: "Um animal se move com  2  pernas."
	animal.Move()

	dog := Cachorro{
		Animal: Animal{Age: 5, Leg: 4},
	}
	//  Output: "Um cachorro se cove com  0  pernas."
	dog.Move()

	dog2 := Cachorro{}
	dog2.Age = 5
	dog2.Leg = 2
	//  Output: "Um cachorro se cove com  0  pernas."
	dog2.Move()

}
