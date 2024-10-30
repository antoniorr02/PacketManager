package internal

type EstadoPedido string

const (
	Pendiente  EstadoPedido = "Pendiente"
	Confirmado              = "Confirmado"
	Entregado               = "Entregado"
	Cancelado               = "Cancelado"
)

type Pedido struct {
	cliente   Usuario
	encargado Usuario
	Pedido    map[TipoPalet]int
	Estado    EstadoPedido
	Id        int
	Factura   float64
}
