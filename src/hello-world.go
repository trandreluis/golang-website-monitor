package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delayEmSegundos = 5

func main() {
	exibeIntroducao()
	for {
		exbibeMenu()

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
	fmt.Println()
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
	fmt.Println()

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		fmt.Println()
		time.Sleep(delayEmSegundos * time.Second)
		fmt.Println()
	}
}

func testaSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Http Status Code:", response.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	arquivo, erroAoAbrir := os.Open("sites.txt")

	if erroAoAbrir != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", erroAoAbrir)
	}

	sites := []string{}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, erroAoLer := leitor.ReadString('\n')
		if erroAoLer == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, erroAoAbrir := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if erroAoAbrir != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo:", erroAoAbrir)
	}

	arquivo.WriteString(site + "- online:" + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
