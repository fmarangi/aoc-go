package day07

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardValues = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardValuesWithJoker = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type Hand struct {
	Cards     string
	Bid, Type int
}

func HandType(cards string) int {
	score := make(map[rune]int, 5)
	for _, c := range cards {
		score[c]++
	}

	var highest int
	for _, n := range score {
		if n > highest {
			highest = n
		}
	}

	switch highest {
	case 5:
		return FiveOfAKind

	case 4:
		return FourOfAKind

	case 3:
		if len(score) == 2 {
			return FullHouse
		}
		return ThreeOfAKind

	case 2:
		if len(score) == 3 {
			return TwoPair
		}
		return OnePair

	}

	return HighCard
}

func HandTypeWithJoker(cards string) int {
	score := make(map[rune]int, 5)
	for _, c := range cards {
		score[c]++
	}

	var highest int
	for c, n := range score {
		if c != 'J' && n > highest {
			highest = n
		}
	}

	highest += score['J']
	delete(score, 'J')

	switch highest {
	case 5:
		return FiveOfAKind

	case 4:
		return FourOfAKind

	case 3:
		if len(score) == 2 {
			return FullHouse
		}
		return ThreeOfAKind

	case 2:
		if len(score) == 3 {
			return TwoPair
		}
		return OnePair

	}

	return HighCard
}
