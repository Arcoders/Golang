package main

import "fmt"

type User interface {
    Permisos() int
    Nombre() string
}

type Admin struct {
    nombre string
}

func (this Admin) Permisos() int {
    return 5
}

func (this Admin) Nombre() string {
    return this.nombre
}

func auth(user User) string {
    permisos := user.Permisos()
    if permisos == 5 {
        return user.Nombre() + " es el amo y lo controla todo..."
    }
    return ""
}

func main() {
    admin := Admin{"Ismael Haytam"}
    usuarios := []User{admin}

    for _,usuario := range usuarios {
        fmt.Println(auth(usuario))
    }
}
