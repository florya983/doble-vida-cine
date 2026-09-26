package main

import(
	"fmt"
)

type Funcion struct{
	pelicula Pelicula
	horario string 
	sala Sala
}

func (f Funcion) VerificarCapacidadSala(entradas int) bool{
	if f.sala.capacidad<entradas{
		fmt.Println("Sala llena")
		return false
		
	}
	return true

}
func (f* Funcion) DescontarCapacidad(entradas int){
	f.sala.capacidad=f.sala.capacidad-entradas 
}

func (f Funcion)VerificarExistenciaPeliculaHora(nombrePelicula string,hora string)bool{
	if diccionarioPeliculas[nombrePelicula]{
		if f.pelicula.nombre==nombrePelicula{
		return f.horario==hora
		}	
	}
	return false
}
