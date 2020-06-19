package main

import "fmt"

type person struct{}

func (p person) Talk() {
	fmt.Println("ငါသည် မုချသေရမည် အချိန်ပိုင်းသာလိုတော့သည်")
}

// Persona လို့ခေါ်တဲ့ symbol ကို export လုပ်လိုက်
var Persona person
