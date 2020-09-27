package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")

	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")

	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }
	// exibeIntroducao()

	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com"
	sites[1] = "https://alura.com.br"
	sites[2] = "https://google.com.br"
	fmt.Println(reflect.TypeOf(sites))
	fmt.Println(sites)
	exibeNomes()

	for {
		// exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			sairDoPrograma()
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Tatiana"
	versao := 1.5
	fmt.Println("Hello, ", nome, "!")
	fmt.Println("Este programa está na versão: ", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi: ", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com"
	sites[1] = "https://alura.com.br"
	sites[2] = "https://google.com.br"
	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, " está com problemas. Status Code: ", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Tatiana", "Anderson", "Luiz"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem: ", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para: ", cap(nomes), "itens")

	nomes = append(nomes, "Moreno")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem: ", len(nomes))
	fmt.Println("O meu slice tem capacidade para: ", cap(nomes), "itens")
}

func exibirLogs() {
	fmt.Println("Exibindo logs...")
}

func sairDoPrograma() {
	fmt.Println("Saindo do programa")
	os.Exit(0)
}
