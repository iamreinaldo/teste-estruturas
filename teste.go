package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func verApp() {
	versao := 0.1
	fmt.Println("Versão: ", versao)
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func mainMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func monitoraMenu() {
	fmt.Println("1 - Genipapo/ Malhadinha/ Balança")
	fmt.Println("2 - OLTs")
	fmt.Println("3 - UNM")
	fmt.Println("Qualquer tecla para voltar ao menu iniciar")
}

func iniciarMonitoramento() {
	fmt.Println("Pingando...")
	enderecos := []string{"172.16.201.2", "172.16.201.98"}
	for _, enderecos := range enderecos {
		ping(enderecos)
	}
}

func ping(enderecos string) {
	out, err := exec.Command("ping", "-c", "4", enderecos).Output()
	if err != nil {
		fmt.Printf("Falha ao executar ping para %s: %s\n\n", enderecos, err)
		return
	}
	fmt.Printf("Resultado do ping para %s:\n%s\n\n", enderecos, strings.TrimSpace(string(out)))
}

func main() {
	verApp()
	mainMenu()
	comando := lerComando()
	switch comando {
	case 1:
		monitoraMenu()
		comandoMonitora := lerComando()
		switch comandoMonitora {
		case 1:
			iniciarMonitoramento()
		case 2:
			iniciarMonitoramento()
		case 3:
			iniciarMonitoramento()
		default:
			iniciarMonitoramento()
		}
	case 2:
		fmt.Println("Imprimindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Print("Comando inválido")
		os.Exit(-1)
	}
}
