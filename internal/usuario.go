package internal

type TipoUsuario string

const (
	Cliente       TipoUsuario = "cliente"
	Encargado                 = "encargado"
	Administrador             = "administrador"
	Transportista             = "transportista"
)

type Usuario struct {
	id       int
	Telefono string
	Rol      TipoUsuario
}
