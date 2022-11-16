package main

import (
	"fmt"
	"net/http"

	//"net/http"
	"os"
)

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
		case 8:
			fmt.Println("Conteúdo de teste iniciado!........")
		case 9:
			fmt.Println("conteudo em teste1.....")
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
	fmt.Println("8- iniciar teste")
	fmt.Println("9- mostrar conteudo de teste")
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

	sites := []string{"https://random-status-code.herokuapp.com/", "www.alura.com.br", "https://www.youtube.com/"}
	//fmt.Println(sites)

	for i, site := range sites {
		fmt.Println("Passando na posição", i, "e está", site)
	}

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "está fora do ar, STATUS:", resp.StatusCode)
	}
}
