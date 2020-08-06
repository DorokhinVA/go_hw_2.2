package card

type Service struct {
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(issuer string, currency string) *Card {
	card := &Card{
		Issuer:   issuer,
		Balance:  0,
		Currency: currency,
		Number:   "0000 0000 0000 0000",
		Icon:     "https://...",
	}
	s.Cards = append(s.Cards, card)
	return card
}

func (s *Service) SearchByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}

func (s *Service) Sum() int64 {
	total := int64(0)
	for _, card := range s.Cards {
		total += card.Balance
	}
	return total
}
