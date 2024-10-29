package internal

type TipoUsuario string

const (
	Cliente       TipoUsuario = "cliente"
	Encargado                 = "encargado"
	Administrador             = "administrador"
	Transportista             = "transportista"
)

type Usuario struct {
	Nombre   string      // Nombre del usuario
	dni      string      // DNI como identificador único del usuario
	Telefono string      // Número de teléfono del usuario
	Rol      TipoUsuario // Rol ya sea cliente, encargado, administrador o transportista.
}
