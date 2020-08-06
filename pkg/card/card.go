package card

func Withdraw(card *Card, amount int64) {
	card.Balance -= amount
}

func Deposit(card *Card, amount int64) {
	card.Balance += amount
}

func Sum(cards []Card) int64 {
	total := int64(0)
	for _, card := range cards {
		total += card.Balance
	}
	return total
}
