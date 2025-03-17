package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

const SHIFT_KEY = 5

func decrypt(message string) string {
    decrypted := ""
    for _, char := range message {
        if char >= 'a' && char <= 'z' {
            decrypted += string((char-'a'-SHIFT_KEY+26)%26 + 'a')
        } else if char >= 'A' && char <= 'Z' {
            decrypted += string((char-'A'-SHIFT_KEY+26)%26 + 'A')
        } else {
            decrypted += string(char)
        }
    }
    return decrypted
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading:", err)
            break
        }

        encryptedMsg := string(message)
        decryptedMsg := decrypt(encryptedMsg)

        fmt.Printf("Mensagem criptografada recebida: %s\n", encryptedMsg)
        fmt.Printf("Mensagem descriptografada: %s\n", decryptedMsg)
    }
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    fmt.Println("Servidor iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
