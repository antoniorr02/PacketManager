package main

import "fmt"

// Definición de la entidad Usuario con diferentes roles
type Usuario struct {
	Nombre   string // Nombre del usuario
	DNI      string // DNI como identificador único del usuario
	Telefono string // Número de teléfono del usuario
	Rol      string // Rol ya sea usuario, encargado o administrador
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

// Pedido representa un pedido de un cliente, gestionado por un encargado
type Pedido struct {
	Cliente   *Usuario
	Encargado *Usuario // Se deberá de comprobar que realmente se trata de un encargado al crear el pedido.
	Sede      *Sede
	Palets    []*Palet
	Estado    string // Estados posibles: "Pendiente", "Confirmado", "Entregado", "Cancelado"
}

// Método para verificar si el usuario es un encargado
func (u *Usuario) EsEncargado() bool {
	return u.Rol == "encargado"
}

// Método para verificar la disponibilidad de palets en una sede
func (s *Sede) VerificarDisponibilidad(tipo string, cantidad int) bool {
	if palet, ok := s.Stock[tipo]; ok && palet.Cantidad >= cantidad {
		return true
	}
	return false
}

// Método para verificar si todos los palets están disponibles en la sede
func (p *Pedido) VerificarDisponibilidad() bool {
	for _, palet := range p.Palets {
		if !p.Sede.VerificarDisponibilidad(palet.Tipo, palet.Cantidad) {
			return false
		}
	}
	return true
}

// Método para procesar el pedido y confirmar disponibilidad de palets
func (p *Pedido) ConfirmarPedido() {
	if p.Encargado.EsEncargado() {
		if p.VerificarDisponibilidad() {
			p.Estado = "Confirmado"
			fmt.Printf("Pedido confirmado para el cliente %s por el encargado %s.\n", p.Cliente.Nombre, p.Encargado.Nombre)
		} else {
			p.Estado = "Pendiente"
			fmt.Printf("Stock insuficiente para el pedido de %s.\n", p.Cliente.Nombre)
		}
	} else {
		fmt.Printf("Error: %s no es un encargado y no puede gestionar el pedido.\n", p.Encargado.Nombre)
	}
}

func main() {
	name := "mundo"
	fmt.Println("Hola", name)
}
