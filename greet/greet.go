package greet

var face = ":)"

// English retorna saludo en ingles
func English() string {
	return "Hi " + face
}

// Italian retorna saludo en italiano
func Italian() string {
	return "Ciao " + face
}

// Deutsche retorna saludo en Aleman
func Deutsche() string {
	return "Hallo " + face
}
