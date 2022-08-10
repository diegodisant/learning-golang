package main

import "fmt"

type Animal struct {
	Name string
	Feet int
}

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

func (d *Dog) bark() {
	fmt.Println("wauf, wauf, wauf...")
}

type Cat struct {
	Animal
}

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

func main() {
	fmt.Println("Creating new dog")

	var arlie Dog

	arlie.Name = "Arlington"
	arlie.Feet = 4

	arlie.sayHi()
	arlie.walk()
	arlie.drink()
	arlie.feed()
	arlie.bark()

	fmt.Println()
	fmt.Println("Creating new dog")

	var dobby *Dog = new(Dog)

	dobby.Name = "Dobby"
	dobby.Feet = 4

	dobby.sayHi()
	dobby.walk()
	dobby.drink()
	dobby.feed()
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

	pillu.sayHi()
	pillu.walk()
	pillu.drink()
	pillu.feed()
	pillu.meow()

	fmt.Println()
	fmt.Println("Creating a new cat")

	var jiji *Cat = new(Cat)

	jiji.Name = "Jiji"
	jiji.Feet = 4

	jiji.sayHi()
	jiji.walk()
	jiji.drink()
	jiji.feed()
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
