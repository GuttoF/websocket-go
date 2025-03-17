package words

import (
	"math/rand"
	"time"
)

type Theme struct {
	Name    string
	Words   []string
}

var Themes = []Theme{
	{
		Name: "Animais",
		Words: []string{
			"cachorro", "elefante", "girassol", "hipopotamo", "papagaio",
			"tigre", "urubu", "zebra", "borboleta",
		},
	},
	{
		Name: "Comida",
		Words: []string{
			"abacaxi", "jabuticaba", "kiwi", "limao", "morango",
			"amendoim", "chocolate", "waffle",
		},
	},
	{
		Name: "Objetos",
		Words: []string{
			"bicicleta", "navio", "sapato", "queijo", "rodovia",
			"violino", "xadrez", "yoga",
		},
	},
}

func GetRandomWordAndTheme() (string, string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	theme := Themes[r.Intn(len(Themes))]
	word := theme.Words[r.Intn(len(theme.Words))]
	return word, theme.Name
}
