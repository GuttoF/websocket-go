package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Conectando ao servidor de Jogo da Forca...")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conectado ao servidor!")

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Conexão com o servidor fechada.")
				os.Exit(0)
				return
			}
			fmt.Print(message)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Entre com uma letra ou 'sair' para sair: ")
		if !scanner.Scan() {
			break
		}

		guess := scanner.Text()
		if strings.ToLower(guess) == "sair" {
			break
		}

		fmt.Fprintf(conn, "%s\n", guess)
	}
}
