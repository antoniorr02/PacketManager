package main

import "fmt"

// Definición de la entidad Cliente
type Cliente struct {
	Nombre   string // Nombre del cliente
	DNI      string // DNI como identificador único del cliente
	Telefono string // Número de teléfono del cliente
}

func main() {
	name := "mundo"
	fmt.Println("Hola", name)
}
