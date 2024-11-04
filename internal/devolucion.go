package internal

type Devolucion struct {
	cliente       Usuario
	transportista Usuario
	Devolucion    map[TipoPalet]map[string]int // String marca el estado
	Compensacion  float64
	Confirmacion  *Confirmacion
	PedidoID      int
}

type Confirmacion struct {
	confirmada    bool
	Administrador Usuario
}
