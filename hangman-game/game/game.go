package game

import (
	"net"
	"strings"
	"fmt"
	"bufio"
	"github.com/guttof/websocket-go/hangman-game/words"
)

type GameState struct {
	Word           string
	GuessedLetters  map[string]bool
	AttemptsLeft    int
	GameOver        bool
	Won            bool
	Theme          string
}

func NewGameState(word, theme string) *GameState {
	return &GameState{
		Word:           word,
		Theme:          theme,
		GuessedLetters: make(map[string]bool),
		AttemptsLeft:   6,
		GameOver:       false,
		Won:            false,
	}
}

func SendGameState(conn net.Conn, game *GameState) {
	var displayWord strings.Builder
	for _, letter := range game.Word {
		if game.GuessedLetters[string(letter)] {
			displayWord.WriteString(string(letter) + " ")
		} else {
			displayWord.WriteString("_ ")
		}
	}

	var guessedLetters strings.Builder
	for letter := range game.GuessedLetters {
		guessedLetters.WriteString(letter + " ")
	}

	stateMsg := fmt.Sprintf("\n=== %s ===\nWord: %s\nGuessed letters: %s\nAttempts left: %d\n",
		game.Theme, displayWord.String(), guessedLetters.String(), game.AttemptsLeft)

	stateMsg += drawHangman(game.AttemptsLeft)

	conn.Write([]byte(stateMsg + "\n"))
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	
	word, theme := words.GetRandomWordAndTheme()
	game := NewGameState(word, theme)
	
	conn.Write([]byte(fmt.Sprintf("Tema: %s\n", theme)))
	
	SendGameState(conn, game)
	
	reader := bufio.NewReader(conn)
	for !game.GameOver {
		guess, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}

		guess = strings.TrimSpace(strings.ToLower(guess))
		if len(guess) != 1 {
			sendMessage(conn, "Por favor, entre uma única letra.")
			continue
		}

		if game.GuessedLetters[guess] {
			sendMessage(conn, "Você já adivinhou essa letra.")
			continue
		}

		game.GuessedLetters[guess] = true

		if !strings.Contains(game.Word, guess) {
			game.AttemptsLeft--
			sendMessage(conn, "Letra não está na palavra!")
		} else {
			sendMessage(conn, "Boa tentativa!")
		}

		won := true
		for _, letter := range game.Word {
			if !game.GuessedLetters[string(letter)] {
				won = false
				break
			}
		}

		if won {
			game.GameOver = true
			game.Won = true
		} else if game.AttemptsLeft <= 0 {
			game.GameOver = true
		}

		SendGameState(conn, game)
	}

	if game.Won {
		sendMessage(conn, "Parabéns! Você ganhou!")
	} else {
		sendMessage(conn, fmt.Sprintf("Jogo terminado! A palavra era: %s", game.Word))
	}
}

func sendMessage(conn net.Conn, message string) {
	conn.Write([]byte(message + "\n"))
}
