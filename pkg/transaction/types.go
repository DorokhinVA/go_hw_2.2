package transaction

import (
	"github.com/DorokhinVA/go_hw_2.2/pkg/card"
	"time"
)

type Transaction struct {
	Id     int64
	Card   *card.Card
	Amount int64
	Date   time.Time
	MCC    string
	Status string
	Type   Type
}

type Type int8

const (
	DEPOSIT  Type = 1
	WITHDRAW Type = 2
)
