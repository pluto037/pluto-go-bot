package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (dog *Dog) Speak() string {
	return "dog name is " + dog.name
}

func (cat *Cat) Speak() string {
	return "cat name is " + cat.name
}

func LetSpeak(speaker Speaker) {
	fmt.Sprintln(speaker.Speak())
}

func main() {
	cat := Cat{name: "ss"}
	LetSpeak(&cat)
}
