package transfer

import (
	"github.com/DorokhinVA/go_hw_2.2/pkg/card"
	"github.com/DorokhinVA/go_hw_2.2/pkg/transaction"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc           *card.Service
		TransactionSvc    *transaction.Service
		MainFeePercent    float64
		AnotherFeePercent float64
		MinFee            int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantErr   error
	}{
		{
			name: "main cards positive",
			fields: fields{CardSvc: &card.Service{
				BankName: "Test Bank",
				Cards: []*card.Card{{
					Id:       1,
					Issuer:   "Visa",
					Balance:  1000_00,
					Currency: "RUB",
					Number:   "5106 211",
				}, {
					Id:       2,
					Issuer:   "Master",
					Balance:  100_000_00,
					Currency: "RUB",
					Number:   "5106 212",
				}},
			},
				TransactionSvc:    &transaction.Service{},
				MainFeePercent:    0.5,
				AnotherFeePercent: 1.5,
				MinFee:            10_00},
			args:      args{from: "5106 212", to: "5106 211", amount: 10_000_00},
			wantTotal: 10_050_00,
			wantErr:   nil,
		},
		{
			name: "main cards negative",
			fields: fields{CardSvc: &card.Service{
				BankName: "Test Bank",
				Cards: []*card.Card{{
					Id:       1,
					Issuer:   "Visa",
					Balance:  1000_00,
					Currency: "RUB",
					Number:   "5106 211",
				}, {
					Id:       2,
					Issuer:   "Master",
					Balance:  100_000_00,
					Currency: "RUB",
					Number:   "5106 212",
				}},
			},
				TransactionSvc:    &transaction.Service{},
				MainFeePercent:    0.5,
				AnotherFeePercent: 1.5,
				MinFee:            10_00},
			args:      args{from: "5106 211", to: "5106 212", amount: 10_000_00},
			wantTotal: 10_050_00,
			wantErr:   ErrNotEnoughMoney,
		},
		{
			name: "main to another positive",
			fields: fields{CardSvc: &card.Service{
				BankName: "Test Bank",
				Cards: []*card.Card{{
					Id:       1,
					Issuer:   "Visa",
					Balance:  1000_00,
					Currency: "RUB",
					Number:   "5106 211",
				}, {
					Id:       2,
					Issuer:   "Master",
					Balance:  100_000_00,
					Currency: "RUB",
					Number:   "5106 212",
				}},
			},
				TransactionSvc:    &transaction.Service{},
				MainFeePercent:    0.5,
				AnotherFeePercent: 1.5,
				MinFee:            10_00},
			args:      args{from: "5106 212", to: "213", amount: 10_000_00},
			wantTotal: 10_050_00,
			wantErr:   ErrTargetCardNotFound,
		},
		{
			name: "main to another negative",
			fields: fields{CardSvc: &card.Service{
				BankName: "Test Bank",
				Cards: []*card.Card{{
					Id:       1,
					Issuer:   "Visa",
					Balance:  1000_00,
					Currency: "RUB",
					Number:   "5106 211",
				}, {
					Id:       2,
					Issuer:   "Master",
					Balance:  100_000_00,
					Currency: "RUB",
					Number:   "5106 212",
				}},
			},
				TransactionSvc:    &transaction.Service{},
				MainFeePercent:    0.5,
				AnotherFeePercent: 1.5,
				MinFee:            10_00},
			args:      args{from: "5106 211", to: "213", amount: 10_000_00},
			wantTotal: 10_050_00,
			wantErr:   ErrNotEnoughMoney,
		},
		{
			name: "another to main",
			fields: fields{CardSvc: &card.Service{
				BankName: "Test Bank",
				Cards: []*card.Card{{
					Id:       1,
					Issuer:   "Visa",
					Balance:  1000_00,
					Currency: "RUB",
					Number:   "5106 211",
				}, {
					Id:       2,
					Issuer:   "Master",
					Balance:  100_000_00,
					Currency: "RUB",
					Number:   "5106 212",
				}},
			},
				TransactionSvc:    &transaction.Service{},
				MainFeePercent:    0.5,
				AnotherFeePercent: 1.5,
				MinFee:            10_00},
			args:      args{from: "212", to: "5106 211", amount: 10_000_00},
			wantTotal: 10_150_00,
			wantErr:   ErrSourceCardNotFound,
		},
		{
			name: "another to another",
			fields: fields{CardSvc: &card.Service{
				BankName: "Test Bank",
				Cards: []*card.Card{{
					Id:       1,
					Issuer:   "Visa",
					Balance:  1000_00,
					Currency: "RUB",
					Number:   "5106 211",
				}, {
					Id:       2,
					Issuer:   "Master",
					Balance:  100_000_00,
					Currency: "RUB",
					Number:   "5106 212",
				}},
			},
				TransactionSvc:    &transaction.Service{},
				MainFeePercent:    0.5,
				AnotherFeePercent: 1.5,
				MinFee:            10_00},
			args:      args{from: "212", to: "213", amount: 10_000_00},
			wantTotal: 10_150_00,
			wantErr:   ErrTargetCardNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:           tt.fields.CardSvc,
				TransactionSvc:    tt.fields.TransactionSvc,
				MainFeePercent:    tt.fields.MainFeePercent,
				AnotherFeePercent: tt.fields.AnotherFeePercent,
				MinFee:            tt.fields.MinFee,
			}
			gotTotal, gotErr := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotErr != tt.wantErr {
				t.Errorf("Card2Card() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
