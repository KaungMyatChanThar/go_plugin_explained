package main

import (
	"fmt"
	"os"
	"plugin"
)

// Persona က စကားတော့ပြောတတ်ရမယ်
type Persona interface {
	Talk()
}

func main() {
	persona := "dev"
	if len(os.Args) == 2 {
		// ရှင်ဘယ်သူလဲ
		persona = os.Args[1]
	}

	mod := fmt.Sprintf("./%[1]v/%[1]v.so", persona)

	// so file ကိုဖွင့်ပြီနော်
	plug, err := plugin.Open(mod)
	if err != nil {
		handleError(err)
	}

	// symbol ကိုရှာ ခုနက export လုပ်ခဲ့တယ်လေ
	symbol, err := plug.Lookup("Persona")
	if err != nil {
		handleError(err)
	}

	// အဆင်ကောပြေရဲ့လား
	var person Persona
	person, ok := symbol.(Persona)
	if !ok {
		handleError(fmt.Errorf("ငါရှာနေတာ ဒါမဟုတ်သေးဘူး"))
	}

	// ပြောကြည့်လေ
	person.Talk()
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
