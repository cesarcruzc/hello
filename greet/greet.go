package greet

// esta vaiable se puede utilizar en cualquier lugar del paquete
// no se puede declarar con el operador de asignacion corta :=
// no permite usar desde otros paquetes cuando inicia con minuscula
// si permite usar desde otros paquetes cuando inicia con mayuscula
var emoji = ":)"

// English retorna saludo en ingles
func English() string {
	return "Hi " + emoji
}

// Italian retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
