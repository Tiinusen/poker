package poker

// Deck represents a deck of cards with basic shuffle and pull functionality
type Deck []Card

// Reset resets deck to a shuffled state with each of every rank and suit
func (d *Deck) Reset(randomizer Randomizer, suits []CardSuit, ranks []CardRank) {
	cards := make(Deck, len(suits)*len(ranks))
	i := 0
	for s := 0; s < len(suits); s++ {
		for r := 0; r < len(ranks); r++ {
			cards[i] = Card{
				CardRank: ranks[r],
				CardSuit: suits[s],
			}
			i++
		}
	}
	randomizer.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	d = &cards
}

// Randomizer is used by a deck to randomize outcomes of shuffles
type Randomizer interface {
	// Shuffle pseudo-randomizes the order of elements using the default Source.
	// n is the number of elements. Shuffle panics if n < 0.
	// swap swaps the elements with indexes i and j.
	Shuffle(n int, swap func(i, j int))
}

// Draw draws amount of cards out of deck from top and returns the altered deck + drawed
func (d Deck) Draw(amount int) (Deck, []Card) {
	if len(d) < amount {
		return Deck{}, d
	}
	return d[amount:], d[:amount]
}
