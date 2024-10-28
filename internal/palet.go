package internal

// Distintos tipos de palets disponibles
type TipoPalet int

const (
	Palet100x100Barra TipoPalet = iota
	PaletPerimetral114x114
	PaletPerimetralFuerte120x100
	PaletPerimetralFino120x100
	PaletCP9
	PaletCP7
	PaletCP6
	PaletCP4
	PaletCP3
	PaletCP1
	Palet80x60Dusseldorf
	Palet90x90Perimetral
	Palet120x100AbiertoSemiFuerte
	PaletAmericano120x100AbiertoFuerteConVuelo
	PaletAmericano120x100AbiertoExtrafuerte
	PaletAmericano120x100AbiertoFuerte
	PaletAmericanoPatinLargo120
	PaletAmericanoPatinCorto100
	Palet120x80Fino
	Palet120x80Fuerte
	EuroPaletHomologadoEPAL
	EuroPaletHomologado1aExtra
)
