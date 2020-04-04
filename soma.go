package main

import "strconv"

func main() {
	println(impressao())
}

func soma(x, y int) int {
	return x + y
}

func impressao() string {
	converte := strconv.Itoa(soma(5, 5))
	resultado := "O resultado da da soma é " + converte
	return resultado
}
