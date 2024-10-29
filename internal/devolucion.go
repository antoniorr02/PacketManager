package internal

type EstadoDevolucion int

const (
	Bueno EstadoDevolucion = iota
	Dañado
	Irreparable
)

// Devolución representa una devolución realizada por un cliente
type Devolucion struct {
	cliente       Usuario
	transportista Usuario
	Devolucion    map[TipoPalet]map[EstadoDevolucion]int
	Compensacion  float64
	Confirmada    bool // Si la devolución fue confirmada por el administrativo
	PedidoID      int  // Asociar devolución con un pedido específico
}
