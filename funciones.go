
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
func (f *Funcion) DescontarCapacidad(entradas int, numeroSala int)bool{
	resultado:=0
	if !ValidarCantidadEntradas(entradas){
		return false
	}
	if !f.ValidarNumeroSala(numeroSala){
		return false
	}
	resultado=f.sala[numeroSala-1] - entradas
	if resultado<0{
		return false
	}
	f.sala[numeroSala-1]=resultado
	return true
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
	if !ValidarCantidadEntradas(entradasPedidas) || !f.ValidarNumeroSala(numeroSala){
		return false
	}
	if f.VerificarExistenciaPeliculaHora(peli.nombre, hora) && f.VerificarCapacidadSala(entradasPedidas, numeroSala) {
		return f.DescontarCapacidad(entradasPedidas, numeroSala)
		
	}
	return false
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
