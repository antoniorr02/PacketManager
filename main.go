package main

import "fmt"

// Definición de la entidad Cliente
type Cliente struct {
	Nombre   string // Nombre del cliente
	DNI      string // DNI como identificador único del cliente
	Telefono string // Número de teléfono del cliente
}

// Palet representa un tipo de palet con una cantidad específica
type Palet struct {
	Tipo     string
	Cantidad int
}

// Sede representa una sede de UPalet con un stock de palets
type Sede struct {
	ID     int               // ID único de la sede
	Nombre string            // Nombre de la sede
	Stock  map[string]*Palet // Mapa de stock de palets, con el tipo como clave
}

func main() {
	name := "mundo"
	fmt.Println("Hola", name)
}
