
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
	if !ValidarCantidadEntradas(entradas){
		return false
	}
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
	if !ValidarCantidadEntradas(entradas){
		return
	}
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
	if !ValidarCantidadEntradas(entradasPedidas){
		return false
	}
	if !f.ValidarNumeroSala(numeroSala){
		return false
	}
	if f.VerificarExistenciaPeliculaHora(peli.nombre, hora) && f.VerificarCapacidadSala(entradasPedidas, numeroSala) {
		f.DescontarCapacidad(entradasPedidas, numeroSala)
	}else {
		return false

	}
	return true 
}
func (f Funcion) ValidarNumeroSala(numeroSala int)bool{
	if !(numeroSala>0 && numeroSala<=len(f.sala)){
		return false
	}
	return true
}
func ValidarCantidadEntradas(entradas int)bool{
	return entradas>0 
}
