package main

import (
	"fmt"
	"github.com/DorokhinVA/go_hw_2.2/pkg/card"
	"github.com/DorokhinVA/go_hw_2.2/pkg/transaction"
	"github.com/DorokhinVA/go_hw_2.2/pkg/transfer"
)

func main() {
	svc := transfer.Service{
		CardSvc: &card.Service{
			BankName: "Test Bank",
			Cards: []*card.Card{{
				Id:       1,
				Issuer:   "Visa",
				Balance:  1000_00,
				Currency: "RUB",
				Number:   "1",
			}, {
				Id:       2,
				Issuer:   "Master",
				Balance:  100_000_00,
				Currency: "RUB",
				Number:   "2",
			}},
		},
		TransactionSvc:    &transaction.Service{},
		MainFeePercent:    0.5,
		AnotherFeePercent: 1.5,
		MinFee:            10_00,
	}

	svc.Card2Card("2", "1", 10_000_00)

	for _, t := range svc.TransactionSvc.Transactions {
		fmt.Println("Transaction: ", t)
	}
}
