package main

import "fmt"

type Animal struct {
	Name string
	Feet int
}

type AnimalInteraction interface {
	sayHi()
	walk()
	feed()
	drink()
}

// force struct to implement the interface
var _ AnimalInteraction = (*Animal)(nil)

func (a *Animal) sayHi() {
	fmt.Println("Hello human my name is", a.Name)
}

func (a *Animal) walk() {
	fmt.Println("Walking with", a.Feet, "foot")
}

func (a *Animal) feed() {
	fmt.Println("crush, crush, crush...")
}

func (a *Animal) drink() {
	fmt.Println("splash, splash, splash...")
}

type Dog struct {
	Animal
}

type DogInteraction interface {
	bark()
}

var _ DogInteraction = (*Dog)(nil)

func (d *Dog) bark() {
	fmt.Println("wauf, wauf, wauf...")
}

type Cat struct {
	Animal
}

type CatInteraction interface {
	meow()
	fight(to Cat)
}

var _ CatInteraction = (*Cat)(nil)

func (c *Cat) meow() {
	fmt.Println("meow, meow, meow...")
}

func (c *Cat) fight(to Cat) {
	fmt.Println()
	fmt.Println("Kitten battle starts")
	fmt.Println(c.Name, "says grrrrr stupid cat", to.Name)
	fmt.Println(c.Name, "cut", to.Name)
	fmt.Println(c.Name, "wins!!!")
}

func runAllInteractions(interaction AnimalInteraction) {
  interaction.sayHi()
	interaction.walk()
	interaction.drink()
	interaction.feed()
}

func main() {
	fmt.Println("Creating new dog")

	var arlie Dog

	arlie.Name = "Arlington"
	arlie.Feet = 4

	runAllInteractions(&arlie)
	arlie.bark()

	fmt.Println()
	fmt.Println("Creating new dog")

	var dobby *Dog = new(Dog)

	dobby.Name = "Dobby"
	dobby.Feet = 4

	runAllInteractions(dobby)
	dobby.bark()

	fmt.Println()
	fmt.Println("Creating new dog")

	var nala Dog = Dog{}

	nala.Animal.Name = "Nala"
	nala.Animal.Feet = 4

	nala.Animal.sayHi()
	nala.Animal.walk()
	nala.Animal.drink()
	nala.Animal.feed()
	nala.bark()

	fmt.Println()
	fmt.Println("Creating a new cat")

	var pillu Cat

	pillu.Name = "Pillu"
	pillu.Feet = 4

	runAllInteractions(&pillu)
	pillu.meow()

	fmt.Println()
	fmt.Println("Creating a new cat")

	var jiji *Cat = new(Cat)

	jiji.Name = "Jiji"
	jiji.Feet = 4

	runAllInteractions(jiji)
	jiji.meow()

	fmt.Println()
	fmt.Println("Creating a new cat")

	var sata Cat = Cat{}

	sata.Animal.Name = "Satanas"
	sata.Animal.Feet = 4

	sata.Animal.sayHi()
	sata.Animal.walk()
	sata.Animal.drink()
	sata.Animal.feed()
	sata.meow()

	pillu.fight(*jiji)
}
