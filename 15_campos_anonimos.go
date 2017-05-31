package main

import "fmt"

type Human struct {
    name string
}

func (this Human) hablar()  string{
    return "El humano " + this.name + " te saluda..."
}

type Tutor struct {
    Human
    Lang
}

func (this Tutor) hablar()  string{
    return "El Tutor " + this.Human.name + " te saluda..."
}

type Lang struct {
    name string
}

func main()  {
    tutor := Tutor{Human{"Ismael Haytam"}, Lang{"Javascript"}}

    fmt.Println(tutor.Lang.name)
    fmt.Println(tutor.Human.name)
    fmt.Println(tutor.hablar())
}
