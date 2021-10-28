package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {
	exibiNomes()

	var sites [4]string
	sites[0] = "https://www.google.com"
	sites[1] = "https://www.stackoverflow.com"
	sites[2] = "https://www.baeldung.com"
	fmt.Println(sites, "Tipo:", reflect.TypeOf(sites))
	// exibeIntroducao()
	for {
		// exbibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}
	}
}

func devolveNome() (string, int) {
	nome := "André"
	idade := 25

	return nome, idade
}

func exibeIntroducao() {
	nome := "André Luís"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exbibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	var sites [4]string
	sites[0] = "https://www.google.com"
	sites[1] = "https://www.stackoverflow.com"
	sites[2] = "https://www.baeldung.com"

	site := "https://random-status-code.herokuapp.com"
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com problemas. Http Status Code:", response.StatusCode)
	}
}

func exibiNomes() {
	nomes := []string{"André", "Fulano", "Sicrano"}
	fmt.Println(nomes, "Tipo", reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Gomes")

	fmt.Println(nomes, "Tipo", reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
