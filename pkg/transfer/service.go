package transfer

import (
	"errors"
	"fmt"
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

var (
	ErrNotEnoughMoney     = errors.New("not enough money on source card")
	ErrSourceCardNotFound = errors.New("source card not found")
	ErrTargetCardNotFound = errors.New("target card not found")
)

func (s *Service) Card2Card(from, to string, amount int64) (total int64, error error) {
	var fromMain bool
	var toMain bool
	var fee int64

	fromCard, err := s.CardSvc.SearchByNumber(from)
	if err == nil {
		fee = s.calculateFee(amount, true)
		fromMain = true
	} else {
		switch err {
		case card.ErrAnotherCardIssuer:
			fee = s.calculateFee(amount, false)
			fmt.Println("Transfer from another issuer card: " + from)
			error = ErrSourceCardNotFound
		default:
			panic("unknown error")
		}
	}
	toCard, err := s.CardSvc.SearchByNumber(to)
	if err == nil {
		toMain = true
	} else {
		switch err {
		case card.ErrAnotherCardIssuer:
			fmt.Println("Transfer to another issuer card: " + to)
			error = ErrTargetCardNotFound
		default:
			panic("unknown error")
		}
	}

	total = amount + fee

	if fromMain {
		if fromCard.Balance < total {
			return total, ErrNotEnoughMoney
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

	if toMain {
		card.Deposit(toCard, amount)
		s.TransactionSvc.AddTransaction(&transaction.Transaction{
			Card:   toCard,
			Amount: amount,
			Date:   time.Now(),
			Type:   transaction.DEPOSIT,
		})
	}

	return total, error

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
