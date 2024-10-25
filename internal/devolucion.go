package internal

// Devolución representa una devolución realizada por un cliente
type Devolucion struct {
	Cliente       *Usuario
	Transportista *Usuario
	Palets        []*Palet
	Compensacion  float64
	Confirmada    bool // Si la devolución fue confirmada por el administrativo
	PedidoID      int  // Asociar devolución con un pedido específico
}
