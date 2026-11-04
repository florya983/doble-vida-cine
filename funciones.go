
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
	if !f.ValidarNumeroSala(numeroSala){
		return false
	}
	if f.sala[numeroSala-1] < entradas {
		fmt.Println("Sala llena")
		return false

	}
	return true
}

func (f *Funcion) DescontarCapacidad(entradas int, numeroSala int){
	if !f.ValidarNumeroSala(numeroSala){
		return 
	}
	f.sala[numeroSala-1] = f.sala[numeroSala-1] - entradas
}

func (f Funcion) VerificarExistenciaPeliculaHora(nombrePelicula string, hora string) bool {
	if diccionarioPeliculas[nombrePelicula] {
		if f.pelicula.nombre == nombrePelicula {
			return f.horario == hora
		}
	}
	return false
}

func (f *Funcion) ReservaFuncion(peli Pelicula, hora string, numeroSala int, entradasPedidas int) bool {
	valor := true
	if !f.ValidarNumeroSala(numeroSala){
		valor=false
		return valor
	}
	if f.VerificarExistenciaPeliculaHora(peli.nombre, hora) && f.VerificarCapacidadSala(entradasPedidas, numeroSala) {
		f.DescontarCapacidad(entradasPedidas, numeroSala)
	}else {
		valor=false

	}
	return valor
}
func (f Funcion) ValidarNumeroSala(numeroSala int)bool{
	if !(numeroSala>0 && numeroSala<=len(f.sala)){
		return false
	}
	return true
}
