// Examplo composição
package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Gender string
}

func (a *Animal) FazSom(name, gender, perform string) {
	fmt.Printf("%s is a %s and %s", name, gender, perform)
	fmt.Println("")
}

type Gato struct {
	Animal
}

func (c *Gato) FazSom() {
	c.Animal.FazSom(c.Name, c.Gender, "meowing")
}

type Cachorro struct {
	Animal
}

func (d *Cachorro) FazSom() {
	d.Animal.FazSom(d.Name, d.Gender, "barking")
}

func main() {
	gato := &Gato{Animal{Name: "Mimi", Gender: "girl"}}
	cachorro := &Cachorro{Animal{Name: "Bob", Gender: "boy"}}

	gato.FazSom()
	cachorro.FazSom()
}
