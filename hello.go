package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 1
const delay = 5

func main() {

	exibeIntrodução()

	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Logs atuais")
		case 0:
			fmt.Println("logoff efetuado com sucesso")
			os.Exit(0)
		default:
			fmt.Println("Este comando não existe")
			os.Exit(-1)
		}
	}
}

func exibeIntrodução() {
	nome := "Alan"
	var versao float32 = 1.1
	fmt.Println("Olá, senhor", nome)
	fmt.Println("Versão atual:", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento dos Sites")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("o comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Minute)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
