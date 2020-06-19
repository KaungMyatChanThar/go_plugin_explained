package main

import "fmt"

type person struct{}

func (p person) Talk() {
	fmt.Println("မင်္ဂလာပါ လောကကြီးရေ")
}

// Persona လို့ခေါ်တဲ့ symbol ကို export လုပ်လိုက်
var Persona person
