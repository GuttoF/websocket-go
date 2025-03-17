package main

import (
	"fmt"
	"net"
	"github.com/guttof/websocket-go/hangman-game/game"
)

func main() {
	fmt.Println("Iniciando servidor de Jogo da Forca...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Servidor está ouvindo na porta 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go game.HandleConnection(conn)
	}
}
