package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
    "github.com/guttof/websocket-go/tic-tac-toe/game"
)

type Server struct {
    game  *game.GameState
    mutex sync.Mutex
}

func NewServer() *Server {
    return &Server{
        game: game.NewGame(),
    }
}

func (s *Server) handleJoin(w http.ResponseWriter, r *http.Request) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    playerNum := 0
    if s.game.CurrentPlayer == 1 {
        playerNum = 1
    } else {
        playerNum = 2
    }

    json.NewEncoder(w).Encode(playerNum)
}

func (s *Server) handleMove(w http.ResponseWriter, r *http.Request) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var move struct {
        Player int `json:"player"`
        Row    int `json:"row"`
        Col    int `json:"col"`
    }

    if err := json.NewDecoder(r.Body).Decode(&move); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if move.Player != s.game.CurrentPlayer {
        http.Error(w, "Não é sua vez", http.StatusBadRequest)
        return
    }

    if !s.game.MakeMove(move.Player, move.Row, move.Col) {
        http.Error(w, "Jogada inválida", http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(s.game)
}

func (s *Server) handleState(w http.ResponseWriter, r *http.Request) {
    s.mutex.Lock()
    defer s.mutex.Unlock()

    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    json.NewEncoder(w).Encode(s.game)
}

func main() {
    server := NewServer()

    http.HandleFunc("/join", server.handleJoin)
    http.HandleFunc("/move", server.handleMove)
    http.HandleFunc("/state", server.handleState)

    log.Println("Servidor iniciado na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
