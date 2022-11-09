package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Alan"
	idade := 24
	var altura = 1.75
	var versao float32 = 1.1

	fmt.Println("olá senhor", nome, "sua idade é:", idade, "com a altura de:", altura)
	fmt.Println("versão atual:", versao)

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))

}
