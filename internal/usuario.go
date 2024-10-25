package internal

type TipoUsuario string

const (
	Cliente       TipoUsuario = "cliente"
	Encargado     TipoUsuario = "encargado"
	Administrador TipoUsuario = "administrador"
	Transportista TipoUsuario = "transportista"
)

type Usuario struct {
	Nombre   string      // Nombre del usuario
	DNI      string      // DNI como identificador único del usuario
	Telefono string      // Número de teléfono del usuario
	Rol      TipoUsuario // Rol ya sea cliente, encargado, administrador o transportista.
}
