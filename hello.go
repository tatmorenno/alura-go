package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

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
	exibeIntroducao()

	for {
		exibeMenu()
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
	fmt.Println("")
	return comandoLido

}
func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com", "https://alura.com.br", "https://google.com.br"}

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando o site: ", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site foi carregado com sucesso")
	} else {
		fmt.Println("Site: ", site, " está com problemas. Status Code: ", resp.StatusCode)
	}
}

func exibirLogs() {
	fmt.Println("Exibindo logs...")
}

func sairDoPrograma() {
	fmt.Println("Saindo do programa")
	os.Exit(0)
}
