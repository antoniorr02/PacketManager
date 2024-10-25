package internal

// Sede representa una sede de UPalet con un stock de palets
type Sede struct {
	ID     int               // ID único de la sede
	Nombre string            // Nombre de la sede
	Stock  map[string]*Palet // Mapa de stock de palets, con el tipo como clave
}
