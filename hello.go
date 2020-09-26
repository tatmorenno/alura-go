package main

import (
	"fmt"
)

func main() {

	exibeIntroducao()
	leComando()

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")

	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")

	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Não conheço este comando")
	}
}

func exibeIntroducao() {
	nome := "Tatiana"
	versao := 1.5
	fmt.Println("Hello, ", nome, "!")
	fmt.Println("Este programa está na versão: ", versao)
}

func leComando(comando) {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi: ", comando)

}
