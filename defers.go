package main

import (
	"fmt"
	"os"
)

func main() {
	file, err :=os.Create("archivo.txt");

	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	defer file.Close()
	_,e:=file.Write([]byte("Hola mundo"))

	if e != nil {
		fmt.Println("Hubo un error")
		return
	}
}