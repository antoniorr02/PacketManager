package internal

type EstadoPedido string

const (
	Pendiente  EstadoPedido = "Pendiente"
	Confirmado              = "Confirmado"
	Entregado               = "Entregado"
	Cancelado               = "Cancelado"
)

// Pedido representa un pedido de un cliente, gestionado por un encargado
type Pedido struct {
	cliente   Usuario
	encargado Usuario // Se deberá de comprobar que realmente se trata de un encargado al crear el pedido.
	Pedido    map[TipoPalet]int
	Estado    EstadoPedido
	Id        int     // Identificador para posibles devoluciones o seguimiento del pedido.
	Factura   float64 // Coste del pedido
}
