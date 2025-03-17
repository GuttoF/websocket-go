package game

type GameState struct {
    Board         [3][3]string
    CurrentPlayer int
    Winner        string
}

func NewGame() *GameState {
    return &GameState{
        Board:         [3][3]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}},
        CurrentPlayer: 1,
        Winner:        "",
    }
}

func (g *GameState) MakeMove(player, row, col int) bool {
    if row < 0 || row > 2 || col < 0 || col > 2 || g.Board[row][col] != "-" {
        return false
    }

    symbol := "X"
    if player == 2 {
        symbol = "O"
    }

    g.Board[row][col] = symbol
    g.CurrentPlayer = 3 - player
    g.checkWinner()
    return true
}

func (g *GameState) checkWinner() {
    for i := 0; i < 3; i++ {
        if g.Board[i][0] != "-" && g.Board[i][0] == g.Board[i][1] && g.Board[i][1] == g.Board[i][2] {
            g.Winner = g.Board[i][0]
            return
        }
    }

    for i := 0; i < 3; i++ {
        if g.Board[0][i] != "-" && g.Board[0][i] == g.Board[1][i] && g.Board[1][i] == g.Board[2][i] {
            g.Winner = g.Board[0][i]
            return
        }
    }

    if g.Board[0][0] != "-" && g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] {
        g.Winner = g.Board[0][0]
        return
    }
    if g.Board[0][2] != "-" && g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0] {
        g.Winner = g.Board[0][2]
        return
    }

    draw := true
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if g.Board[i][j] == "-" {
                draw = false
            }
        }
    }
    if draw {
        g.Winner = "DRAW"
    }
}
