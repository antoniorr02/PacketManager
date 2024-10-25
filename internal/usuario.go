package internal

// Definición de la entidad Usuario con diferentes roles
type Usuario struct {
	Nombre   string // Nombre del usuario
	DNI      string // DNI como identificador único del usuario
	Telefono string // Número de teléfono del usuario
	Rol      string // Rol ya sea usuario, encargado, administrador o transportista.
}
