package main

import (
	"bufio"
	"fmt"
	"io"
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

func iniciarMonitoramento(nome string) {
	fmt.Println("Pingando...")
	enderecos := lerEnderecosArquivo(nome)
	for _, enderecos := range enderecos {
		ping(enderecos)
	}
}

func lerEnderecosArquivo(nome string) []string {
	var ips []string
	arquivo, err := os.Open(nome)

	if err != nil {
		fmt.Println("Falha ao ler o arquivo: ", err)
	}

	defer arquivo.Close()
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if linha != "" {
			ips = append(ips, linha)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Erro ao ler a linha: ", err)
			break
		}
	}
	fmt.Println("Pingando para os seguintes endereços: ", ips)
	return ips
}

func ping(enderecos string) {
	err := exec.Command("ping", "-c", "4", enderecos).Run()
	if err == nil {
		fmt.Printf("Ping para %s bem-sucedido. \n\n", enderecos)
	} else {
		fmt.Printf("Falha ao executar ping para %s. \n\n", enderecos)
	}
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
			iniciarMonitoramento("balanca.txt")
		case 2:
			iniciarMonitoramento("olt.txt")
		case 3:
			iniciarMonitoramento("unm.txt")
		default:
			os.Exit(-1)
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
