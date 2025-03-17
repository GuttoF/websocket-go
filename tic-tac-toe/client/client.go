package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "github.com/guttof/websocket-go/tic-tac-toe/game"
)

func printBoard(board [3][3]string) {
    fmt.Println("\n  0 1 2")
    for i := 0; i < 3; i++ {
        fmt.Printf("%d ", i)
        for j := 0; j < 3; j++ {
            fmt.Printf("%s ", board[i][j])
        }
        fmt.Println()
    }
    fmt.Println()
}

func getGameState() (*game.GameState, error) {
    resp, err := http.Get("http://localhost:8080/state")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var gameState game.GameState
    if err := json.NewDecoder(resp.Body).Decode(&gameState); err != nil {
        return nil, err
    }

    return &gameState, nil
}

func makeMove(player, row, col int) error {
    move := struct {
        Player int `json:"player"`
        Row    int `json:"row"`
        Col    int `json:"col"`
    }{
        Player: player,
        Row:    row,
        Col:    col,
    }

    moveJSON, err := json.Marshal(move)
    if err != nil {
        return err
    }

    resp, err := http.Post("http://localhost:8080/move", "application/json", bytes.NewBuffer(moveJSON))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("erro ao fazer jogada: %s", resp.Status)
    }

    return nil
}

func main() {
    resp, err := http.Post("http://localhost:8080/join", "application/json", nil)
    if err != nil {
        fmt.Println("Erro ao conectar ao servidor:", err)
        return
    }
    defer resp.Body.Close()

    var playerNum int
    if err := json.NewDecoder(resp.Body).Decode(&playerNum); err != nil {
        fmt.Println("Erro ao decodificar resposta:", err)
        return
    }

    fmt.Printf("Você é o jogador %d (", playerNum)
    if playerNum == 1 {
        fmt.Println("X)")
    } else {
        fmt.Println("O)")
    }

    for {
        gameState, err := getGameState()
        if err != nil {
            fmt.Println("Erro ao obter estado do jogo:", err)
            continue
        }

        printBoard(gameState.Board)

        if gameState.Winner != "" {
            if gameState.Winner == "DRAW" {
                fmt.Println("Jogo empatado!")
            } else {
                fmt.Printf("Jogador %s venceu!\n", gameState.Winner)
            }
            return
        }

        if gameState.CurrentPlayer != playerNum {
            fmt.Println("Aguardando outro jogador...")
            time.Sleep(1 * time.Second)
            continue
        }

        var input string
        fmt.Print("Sua vez! Digite a linha e coluna (0-2): ")
        fmt.Scan(&input)
        
        var row, col int
        _, err = fmt.Sscanf(input, "%d-%d", &row, &col)
        if err != nil {
            fmt.Println("Formato inválido. Use 'linha-coluna' (por exemplo: 0-1)")
            continue
        }


        if err := makeMove(playerNum, row, col); err != nil {
            fmt.Println("Erro:", err)
            continue
        }
    }
}
