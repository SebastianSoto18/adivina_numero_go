package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numero := rand.Intn(100)
	vidas := 10
	acerto := false
	intento := 0
	for !acerto || vidas > 0 {
		fmt.Println("Ingrese un numero: ")
		fmt.Scanln(&intento)

		if intento == numero {
			fmt.Println("Acertaste")
			break
		}else{
			pista(numero, intento)
			vidas--
			fmt.Println("Te quedan ", vidas, " vidas")
		}
	}

}

func pista(numero int, intento int){
	if intento < numero {
		fmt.Println("El numero es mayor")
	}else{
		fmt.Println("El numero es menor")
	}
}