package main

import "fmt"

type User struct {
    edad int
    nombre, apellido string
}

func main() {

    user := User{nombre: "Ismael Haytam", apellido: "Tanan", edad: 21}

    fmt.Println(user)

    fmt.Println(user.nombre)

}
