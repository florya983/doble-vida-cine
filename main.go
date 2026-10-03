package main

import "fmt"

func main() {
	funcion := Funcion{
		pelicula: Pelicula{nombre: "Avatar"},
		horario:  "20:00",
		sala:     listaSalasCapacidad,
	}
	fmt.Println("¿Existe Avatar a las 20:00?")
	fmt.Println(funcion.VerificarExistenciaPeliculaHora("Avatar", "20:00"))

	fmt.Println("¿Hay capacidad para 10 entradas?")
	fmt.Println(funcion.VerificarCapacidadSala(10, 2))

	fmt.Println("Capacidad antes de reservar:")
	fmt.Println(funcion.sala)

	funcion.DescontarCapacidad(10, 0)

	fmt.Println("Capacidad después de reservar 10 entradas:")
	fmt.Println(funcion.sala)

}

