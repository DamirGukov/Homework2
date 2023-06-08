package main

import "fmt"

type Zookeeper struct {
	Name string
}

type Animal struct {
	Name string
}

type Cage struct {
	Cage *Animal
}

func (z *Zookeeper) CollectAnimal(a *Animal, c *Cage) (string, string) {
	c.Cage = a
	return "Zookeeper " + z.Name + " catched  animal " + a.Name, ""
}

func main() {

	z := Zookeeper{Name: "Andrew"}

	animal1 := &Animal{Name: "Tiger"}
	animal2 := &Animal{Name: "Penguin"}
	animal3 := &Animal{Name: "Monkey"}
	animal4 := &Animal{Name: "Horse"}
	animal5 := &Animal{Name: "Lion"}

	cage1 := &Cage{}
	cage2 := &Cage{}
	cage3 := &Cage{}
	cage4 := &Cage{}
	cage5 := &Cage{}

	fmt.Println(z.CollectAnimal(animal1, cage1))
	fmt.Println(z.CollectAnimal(animal2, cage2))
	fmt.Println(z.CollectAnimal(animal3, cage3))
	fmt.Println(z.CollectAnimal(animal4, cage4))
	fmt.Println(z.CollectAnimal(animal5, cage5))

}
