package card

import (
	"errors"
	"strings"
)

type Service struct {
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

const MainCardPrefix = "5106 21"

var (
	ErrAnotherCardIssuer = errors.New("another card issuer")
)

func (s *Service) IssueCard(issuer string, currency string) *Card {
	card := &Card{
		Issuer:   issuer,
		Balance:  0,
		Currency: currency,
		Number:   "5106 2100 0000 0000",
		Icon:     "https://...",
	}
	s.Cards = append(s.Cards, card)
	return card
}

func (s *Service) SearchByNumber(number string) (card *Card, err error) {
	if !s.verifyCardNumber(number) {
		return nil, ErrAnotherCardIssuer
	}

	for _, card := range s.Cards {
		if card.Number == number {
			return card, nil
		}
	}

	panic("couldn`t find card by valid number: " + number)
}

func (s *Service) verifyCardNumber(number string) bool {
	if strings.HasPrefix(number, MainCardPrefix) {
		return true
	}
	return false
}

func (s *Service) Sum() int64 {
	total := int64(0)
	for _, card := range s.Cards {
		total += card.Balance
	}
	return total
}
