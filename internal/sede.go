package internal

// Sede representa una sede de UPalet con un stock de palets
type Sede struct {
	id     int               // ID único de la sede
	nombre string            // Nombre de la sede
	Stock  map[TipoPalet]int // Stock de cada tipo de palet dentro de la sede
}
