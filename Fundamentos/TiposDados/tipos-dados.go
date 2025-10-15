package main

import "errors"

func main() {
	// int8, int16, int32, int64
	// se não especificar o tamanho, o padrão é int32 ou int64 dependendo do sistema operacional
	// int é um tipo de dado que pode armazenar números inteiros, positivos e negativos
	// uint é um tipo de dado que pode armazenar apenas números inteiros positivos

	var numero int32 = 100
	numero2 := int64(numero)
	println(numero2)

	//alias
	// int32 = rune
	// uint8 = byte
	var numero3 rune = 123456
	var numero4 byte = 123
	println(numero4)
	println(numero3)

	// float32, float64
	var numeroReal1 float32 = 123.456
	var numeroReal2 float64 = 123.45678901234
	println(numeroReal1)
	println(numeroReal2)

	// string
	var str string = "Olá, mundo!"
	println(str)
	var str2 string = `Olá	,
	mundo!`
	println(str2)

	//char (na verdade é um int32 que representa o código Unicode)
	// char retorna o valor numérico do caractere na tabela Unicode
	var char rune = 'A'
	println(char)
	char = 'B'
	println(char)
	char = 67 // código Unicode do caractere 'C'

	//para go todo tipo de dado tem o valor inicial padrão
	var str3 string
	println(str3) // valor inicial padrão é vazio
	var inteiro int
	println(inteiro) // valor inicial padrão é 0
	var boolean bool
	println(boolean) // valor inicial padrão é false

	// bollean
	var booleano bool = true
	println(booleano)
	booleano = false
	println(booleano)

	// tipo error (interface)
	var erro error
	println(erro) // valor inicial padrão é nil
	erro = nil
	println(erro)

	var erro2 error = errors.New("Erro interno")
	println(erro2)
	

}