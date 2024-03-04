package main;

import (
	"fmt"
	"math"
)

func main() {
	var (
		base float64
		altura float64
		hipotenusa float64
	)

	fmt.Print("Ingrese la base del triángulo: ")
	fmt.Scan(&base)
	fmt.Print("Ingrese la altura del triángulo: ")
	fmt.Scan(&altura)

	hipotenusa = math.Sqrt(math.Pow(float64(base), 2) + math.Pow(float64(altura), 2))

	fmt.Printf("La hipotenusa del triángulo es: %.2f \n", hipotenusa)
	fmt.Printf("El perimetro del triángulo es: %.2f", base + altura + hipotenusa)
}