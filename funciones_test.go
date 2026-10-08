
package main

import (
	"testing"
)

func TestVerificaCapacidad(t *testing.T) {
	funcion := Funcion{
		sala: listaSalasCapacidad,
	}
	entradas := 15
	numeroSala := 2
	resultado := funcion.VerificarCapacidadSala(entradas, numeroSala)
	esperado := false
	if resultado != esperado {
		t.Errorf("El resultado %v es diferente al esperado %v", resultado, esperado)
	}

}

func TestVerificarExistenciaPelicula(t *testing.T) {
	funcion := Funcion{
		pelicula: Pelicula{
			nombre: "Avatar",
		},
		horario: "20:00",
		sala:    listaSalasCapacidad,
	}
	hora := "20:00"
	nombrePelicula := "Avatar"
	resultado := funcion.VerificarExistenciaPeliculaHora(nombrePelicula, hora)
	esperado := true
	if resultado != esperado {
		t.Errorf("El resultado %v es diferente al esperado %v", nombrePelicula, hora)
	}

}

