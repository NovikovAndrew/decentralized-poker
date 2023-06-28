package deck

import (
	"math/rand"
	"time"
)

const (
	nCards  = 52
	nSuits  = 4
	nValues = 12
)

type Deck [nCards]Card

func NewDeck() Deck {
	var (
		cards = [nCards]Card{}
		x     = 0
	)

	for s := 0; s < nSuits; s++ {
		for v := 0; v < nValues; v++ {
			cards[x] = NewCard(Suit(s), v+1)
			x++
		}
	}

	return shuffle(cards)
}

func shuffle(d Deck) Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	return d
}
