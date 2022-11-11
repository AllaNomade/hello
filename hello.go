package main

import (
	"fmt"
	"net/http"
	"reflect"

	//"net/http"
	"os"
)

func main() {

	exibeNomes()

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

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "www.alura.com.br"
	sites[2] = "www.youtube.com"
	fmt.Println(sites[3])

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso!")
	} else {
		fmt.Println("O site:", site, "está fora do ar, STATUS:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"alan", "caro", "samanta", "vasco"}
	fmt.Println(nomes)

	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("slice tem a capacidade de ", len(nomes))
	nomes = append(nomes, "Aparecido")

	fmt.Println("E AGORA A CAPACIDADE DE:", len(nomes))
	fmt.Println("NOVA CAPACIDADE", cap(nomes))
}
