package deck

import "fmt"

type Card struct {
	Suit  Suit
	Value int
}

func NewCard(s Suit, v int) Card {
	if v > 13 || v < 0 {
		panic("the value of the cards cannot be higher 13 or lower 1")
	}

	return Card{
		Suit:  s,
		Value: v,
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%d%s", c.Value, c.Suit.SuitToUnicode())
}
