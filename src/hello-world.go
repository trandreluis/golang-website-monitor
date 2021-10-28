package main

import (
	"fmt"
)

func main() {
	nome := "André Luís"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variável comando é", &comando)

	if comando == 1 {
		fmt.Println("A opção 1 foi escolhida")
	} else if comando == 2 {
		fmt.Println("A opção 2 foi escolhida")
	} else if comando == 0 {
		fmt.Println("A opção 0 foi escolhida")
	}
}
