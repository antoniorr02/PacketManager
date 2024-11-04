package internal

type Pedido struct {
	cliente   Usuario
	encargado Usuario
	Pedido    map[TipoPalet]int
	Estado    string
	Id        int
	Factura   float64
}
