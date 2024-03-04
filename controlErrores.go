package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0,errors.New("No se puede dividir por cero");
	}
	return num1 / num2, nil
}


func main() {
	str := "123"

	// Convertir un string a un entero
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println(num)
	}

	fmt.Println(divide(10, 0))
	fmt.Println(divide(10, 2))
}