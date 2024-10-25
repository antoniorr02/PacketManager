package internal

// Pedido representa un pedido de un cliente, gestionado por un encargado
type Pedido struct {
	Cliente   *Usuario
	Encargado *Usuario // Se deberá de comprobar que realmente se trata de un encargado al crear el pedido.
	Palets    []*Palet
	Estado    string  // Estados posibles: Pendiente, Confirmado, Entregado, Cancelado.
	Id        int     // Identificador para posibles devoluciones o seguimiento del pedido.
	Factura   float64 // Coste del pedido
}
