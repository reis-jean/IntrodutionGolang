package main

import "fmt"

func main() {
	var variavel string = "Variavel 1" // forma tradicional de declarar uma variavel
	variavel2 := "Variavel 2" // forma simplificada de declarar uma variavel 
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	) // forma de declarar varias variaveis de uma vez 

	variavel5, variavel6 := "Variavel 5", "Variavel 6" // forma de declarar varias variaveis de uma vez 

	fmt.Println(variavel, variavel2, variavel3, variavel4, variavel5, variavel6)

	const constante string = "Constante" // forma de declarar uma constante
	fmt.Println(constante)

	variavel5, variavel6 = variavel6, variavel5 // forma de trocar o valor de duas variaveis
	fmt.Println(variavel5, variavel6)
	
}