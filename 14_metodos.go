package main

import "fmt"

type User struct {
    edad int
    nombre, apellido string
}

func (this User) nombre_completo() string{
    return this.nombre + " " + this.apellido
}

func (this *User) set_name(name string) {
    this.nombre = name
}

func main() {
    user := new(User)

    user.nombre = "Ismael Haytam"
    user.apellido = "Tanan"
    user.edad = 21
    fmt.Println(user.nombre_completo())

    user.set_name("Fatima Chadli")
    fmt.Println(user.nombre)
}
