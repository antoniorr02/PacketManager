package internal

type EstadoDevolucion int

const (
	Bueno EstadoDevolucion = iota
	Dañado
	Irreparable
)

type Devolucion struct {
	cliente       Usuario
	transportista Usuario
	Devolucion    map[TipoPalet]map[EstadoDevolucion]int
	Compensacion  float64
	Confirmacion  *Confirmacion
	PedidoID      int
}

type Confirmacion struct {
	confirmada    bool
	Administrador Usuario
}
