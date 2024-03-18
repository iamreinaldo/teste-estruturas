package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Endereco struct {
	IP   string
	Nome string
}

func verApp() {
	versao := "1.0.1"
	fmt.Println("Versão: ", versao)
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func mainMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	// fmt.Println("2 - Exibir logs")
	fmt.Printf("0 - Sair do programa\n\n")

}

func monitoraMenu() {
	fmt.Println("1 - Genipapo/ Malhadinha/ Balança")
	fmt.Println("2 - OLTs")
	fmt.Println("3 - Lages do Batata")
	fmt.Println("4 - Paraíso")
	fmt.Println("5 - Palmeirinha")
	fmt.Printf("Qualquer tecla para voltar ao menu iniciar\n\n")
}

func iniciarMonitoramento(nome string) {
	fmt.Println("Pingando...")
	enderecos := lerEnderecosArquivo(nome)
	if enderecos == nil {
		fmt.Println("Não há endereços para monitorar")
		return
	}

	for _, enderecos := range enderecos {
		ping(enderecos)
	}
}

func lerEnderecosArquivo(nome string) []Endereco {
	var ips []Endereco
	arquivo, err := os.Open(nome)

	if err != nil {
		fmt.Println("Falha ao ler o arquivo: ", err)
	}

	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if linha != "" {
			partes := strings.Split(linha, ";")
			if len(partes) == 2 {
				ips = append(ips, Endereco{
					IP:   partes[0],
					Nome: partes[1],
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		return nil
	}

	for _, endereco := range ips {
		fmt.Printf("Pingando para %s (%s)\n", endereco.Nome, endereco.IP)
	}
	return ips
}

func ping(enderecos Endereco) {
	err := exec.Command("ping", "-c", "4", enderecos.IP).Run()
	if err == nil {
		fmt.Printf("\nPing para %s bem-sucedido. \n", enderecos.Nome)
	} else {
		fmt.Printf("\nFalha ao executar ping para %s. \n", enderecos.Nome)
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
			iniciarMonitoramento("lages.txt")
		case 4:
			iniciarMonitoramento("paraiso.txt")
		case 5:
			iniciarMonitoramento("palmeirinha.txt")
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
