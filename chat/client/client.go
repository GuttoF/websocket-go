// client.go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "github.com/gorilla/websocket"
)

const SHIFT_KEY = 5

func encrypt(message string) string {
    encrypted := ""
    for _, char := range message {
        if char >= 'a' && char <= 'z' {
            encrypted += string((char-'a'+SHIFT_KEY)%26 + 'a')
        } else if char >= 'A' && char <= 'Z' {
            encrypted += string((char-'A'+SHIFT_KEY)%26 + 'A')
        } else {
            encrypted += string(char)
        }
    }
    return encrypted
}

func main() {
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
    if err != nil {
        log.Fatal("Erro ao conectar:", err)
    }
    defer conn.Close()

    fmt.Println("Conectado ao servidor. Digite suas mensagens:")
    scanner := bufio.NewScanner(os.Stdin)

    for {
        if scanner.Scan() {
            message := scanner.Text()
            encryptedMsg := encrypt(message)

            err := conn.WriteMessage(websocket.TextMessage, []byte(encryptedMsg))
            if err != nil {
                log.Println("Erro ao enviar mensagem:", err)
                return
            }

            fmt.Printf("Mensagem enviada (criptografada): %s\n", encryptedMsg)
        }
    }
}
