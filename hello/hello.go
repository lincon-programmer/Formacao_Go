package main

import "fmt"

func main() {
	nome := "Lincon"
	idade := 23
	versao := 1.1
	fmt.Println("Olá, sr.", nome, "Sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O endereço da minha variável comando é", &comando)
	fmt.Println("O valor escolhido foi", comando)
}
