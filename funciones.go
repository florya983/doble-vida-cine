package main

import (
	"fmt"
)

type Funcion struct {
	pelicula Pelicula
	horario  string
	sala     []int
}

func (f Funcion) VerificarCapacidadSala(entradas int, numeroSala int) bool {
	if f.sala[numeroSala] < entradas {
		fmt.Println("Sala llena")
		return false

	}
	return true

}
func (f *Funcion) DescontarCapacidad(entradas int, numeroSala int) {
	f.sala[numeroSala] = f.sala[numeroSala] - entradas
}

func (f Funcion) VerificarExistenciaPeliculaHora(nombrePelicula string, hora string) bool {
	if diccionarioPeliculas[nombrePelicula] {
		if f.pelicula.nombre == nombrePelicula {
			return f.horario == hora
		}
	}
	return false
}

func (f Funcion) ReservaFuncion(peli Pelicula, hora string, numeroSala int, entradasPedidas int) bool {
	valor := true
	if !f.VerificarExistenciaPeliculaHora(peli.nombre, hora) {
		valor = false
	}
	if !f.VerificarCapacidadSala(entradasPedidas, numeroSala) {
		valor = false
	} else {
		f.DescontarCapacidad(entradasPedidas, numeroSala)

	}
	return valor
}

