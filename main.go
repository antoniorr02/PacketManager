package main

import "fmt"

// Definición de la entidad Usuario con diferentes roles
type Usuario struct {
	Nombre   string // Nombre del usuario
	DNI      string // DNI como identificador único del usuario
	Telefono string // Número de teléfono del usuario
	Rol      string // Rol ya sea usuario, encargado, administrador o transportista.
}

// Palet representa un tipo de palet con una cantidad específica
type Palet struct {
	Tipo     string
	Cantidad int
	Estado   string // Estado del palet
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
	Palets    []*Palet
	Estado    string  // Estados posibles: Pendiente, Confirmado, Entregado, Cancelado.
	Id        int     // Identificador para posibles devoluciones o seguimiento del pedido.
	Factura   float64 // Coste del pedido
}

// Devolución representa una devolución realizada por un cliente
type Devolucion struct {
	Cliente       *Usuario
	Transportista *Usuario
	Palets        []*Palet
	Compensacion  float64
	Confirmada    bool // Si la devolución fue confirmada por el administrativo
	PedidoID      int  // Asociar devolución con un pedido específico
}

// Método para verificar si el usuario es un encargado
func (u *Usuario) EsEncargado() bool {
	return u.Rol == "encargado"
}

// Método para verificar la disponibilidad de palets en una sede
/*
	La idea es hacer un bucle que recorra todos los productos de un pedido.
	Dentro del primer bucle recorrer todas las sedes con otro bucle.
	De esta forma se confirma que haya stock de todos los productos del pedido en alguna de las sedes.
*/
func (s *Sede) VerificarDisponibilidad(tipo string, cantidad int) bool {
	if palet, ok := s.Stock[tipo]; ok && palet.Cantidad >= cantidad {
		return true
	}
	return false
}

// Método para procesar el pedido y confirmar disponibilidad de palets
func (p *Pedido) ConfirmarPedido() {
	if p.Encargado.EsEncargado() {
		p.Estado = "Confirmado"
		fmt.Printf("Pedido confirmado para el cliente %s por el encargado %s.\n", p.Cliente.Nombre, p.Encargado.Nombre)
	} else {
		fmt.Printf("Error: %s no es un encargado y no puede gestionar el pedido.\n", p.Encargado.Nombre)
	}
}

// Método para validar los palets devueltos y registrar la devolución
func (d *Devolucion) ValidarDevolucion() {
	fmt.Printf("Validando la devolución realizada por %s:\n", d.Cliente.Nombre)
	totalCompensacion := 0.0
	for _, palet := range d.Palets {
		fmt.Printf("- %d palets de tipo %s (%s)\n", palet.Cantidad, palet.Tipo, palet.Estado)
		switch palet.Estado {
		case "nuevo":
			totalCompensacion += float64(palet.Cantidad) * 10.0 // Ejemplo de compensación por palet nuevo
		case "usado":
			totalCompensacion += float64(palet.Cantidad) * 5.0 // Compensación por palet usado
		case "dañado":
			// No se agrega al stock si está dañado, solo se registra la merma
			fmt.Println("Palet dañado, no se devuelve al stock.")
		}
	}
	d.Compensacion = totalCompensacion
	fmt.Printf("Compensación total para %s: %.2f euros\n", d.Cliente.Nombre, d.Compensacion)
}

// Administrativo confirma la devolución y ajusta la facturación del cliente
func ConfirmarDevolucion(devolucion *Devolucion, pedido *Pedido, administrativo *Usuario) {
	if administrativo.Rol != "administrativo" {
		fmt.Println("Error: solo un administrativo puede confirmar la devolución.")
		return
	}

	if !devolucion.Confirmada {
		devolucion.Confirmada = true
		pedido.Factura -= devolucion.Compensacion // Ajustar la factura con la compensación
		fmt.Printf("Devolución confirmada por %s. Factura del pedido ajustada en %.2f euros.\n", administrativo.Nombre, devolucion.Compensacion)
		fmt.Printf("Nueva factura del pedido: %.2f euros\n", pedido.Factura)
	} else {
		fmt.Println("La devolución ya ha sido confirmada anteriormente.")
	}
}

func main() {
	// Creación de usuarios
	cliente := &Usuario{Nombre: "Juan Pérez", DNI: "12345678A", Telefono: "555-1234", Rol: "usuario"}
	encargado := &Usuario{Nombre: "Ana Gómez", DNI: "87654321B", Telefono: "555-5678", Rol: "encargado"}
	administrativo := &Usuario{Nombre: "Luis López", DNI: "11223344C", Telefono: "555-9101", Rol: "administrativo"}
	transportista := &Usuario{Nombre: "Carlos Ruiz", DNI: "99887766D", Telefono: "555-1213", Rol: "transportista"}

	// Creación de sedes
	sede := &Sede{
		ID:     1,
		Nombre: "Sede Central",
		Stock: map[string]*Palet{
			"palet_nuevo": {Tipo: "palet_nuevo", Cantidad: 20, Estado: "nuevo"},
		},
	}

	// Verificación de disponibilidad de palets
	tipoPalet := "palet_nuevo"
	cantidad := 5
	if sede.VerificarDisponibilidad(tipoPalet, cantidad) {
		fmt.Printf("Disponibilidad confirmada para %d palets de tipo %s en %s.\n", cantidad, tipoPalet, sede.Nombre)
	} else {
		fmt.Printf("No hay suficiente stock de %s en %s.\n", tipoPalet, sede.Nombre)
	}

	// Creación de un pedido
	pedido := &Pedido{
		Cliente:   cliente,
		Encargado: encargado,
		Palets:    []*Palet{{Tipo: "palet_nuevo", Cantidad: 5, Estado: "nuevo"}},
		Estado:    "Pendiente",
		Id:        1,
		Factura:   50.0,
	}

	pedido.ConfirmarPedido()

	// Validación de la devolución
	devolucion := &Devolucion{
		Cliente:       cliente,
		Transportista: transportista,
		Palets: []*Palet{
			{Tipo: "palet_nuevo", Cantidad: 2, Estado: "nuevo"},
			{Tipo: "palet_usado", Cantidad: 1, Estado: "usado"},
			{Tipo: "palet_dañado", Cantidad: 1, Estado: "dañado"},
		},
		PedidoID: 1,
	}

	// Validar y mostrar compensación de la devolución
	devolucion.ValidarDevolucion()

	// Confirmar la devolución
	ConfirmarDevolucion(devolucion, pedido, administrativo)
}
