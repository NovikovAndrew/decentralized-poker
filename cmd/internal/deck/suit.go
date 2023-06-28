package deck

type Suit int

const (
	Spades Suit = iota
	Harts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "Spades"
	case Harts:
		return "Harts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	default:
		panic("invalid card")
	}
}

func (s Suit) SuitToUnicode() string {
	switch s {
	case Spades:
		return "♠"
	case Harts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("invalid card")
	}
}
