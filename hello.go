package main

import (
	"fmt"
	"net/http"

	//"net/http"
	"os"
)

func main() {

	exibeIntrodução()
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

func exibeIntrodução() {
	nome := "Alan"
	var versao float32 = 1.1
	fmt.Println("Olá, senhor", nome)
	fmt.Println("Versão atual:", versao)
}

func exibeRedbull() (string, int) {
	equipe := "Redbull"
	pilotos := 1

	return equipe, pilotos
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

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://app.awari.com.br/"
	resp, _ := http.Get(site)
	fmt.Println(resp)

}
