package transfer

import (
	"github.com/DorokhinVA/go_hw_2.2/pkg/card"
	"github.com/DorokhinVA/go_hw_2.2/pkg/transaction"
	"math"
	"time"
)

type Service struct {
	CardSvc           *card.Service
	TransactionSvc    *transaction.Service
	MainFeePercent    float64
	AnotherFeePercent float64
	MinFee            int64
}

func NewService(cardSvc *card.Service, mainFeePercent float64, anotherFeePercent float64, minFee int64) *Service {
	return &Service{CardSvc: cardSvc, MainFeePercent: mainFeePercent, AnotherFeePercent: anotherFeePercent, MinFee: minFee}
}

func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {
	fromCard := s.CardSvc.SearchByNumber(from)
	toCard := s.CardSvc.SearchByNumber(to)

	var fee int64
	if fromCard == nil {
		fee = s.calculateFee(amount, false)
	} else {
		fee = s.calculateFee(amount, true)
	}
	total = amount + fee

	if fromCard != nil {
		if fromCard.Balance < total {
			return total, ok
		}

		card.Withdraw(fromCard, total)
		s.TransactionSvc.AddTransaction(&transaction.Transaction{
			Card:   fromCard,
			Amount: amount,
			Date:   time.Now(),
			Type:   transaction.WITHDRAW,
		})
		if fee > 0 {
			s.TransactionSvc.AddTransaction(&transaction.Transaction{
				Card:   fromCard,
				Amount: fee,
				Date:   time.Now(),
				Type:   transaction.WITHDRAW,
			})
		}
	}

	if toCard != nil {
		card.Deposit(toCard, amount)
		s.TransactionSvc.AddTransaction(&transaction.Transaction{
			Card:   toCard,
			Amount: amount,
			Date:   time.Now(),
			Type:   transaction.DEPOSIT,
		})
	}

	return total, true

}

func (s *Service) calculateFee(amount int64, main bool) int64 {
	var percent float64
	if main {
		percent = s.MainFeePercent
	} else {
		percent = s.AnotherFeePercent
	}

	fee := float64(amount) / 100 * percent

	if math.Round(fee) < float64(s.MinFee) {
		return s.MinFee
	}

	return int64(fee)
}
