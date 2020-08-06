package transaction

type Service struct {
	Transactions []*Transaction

	Index int64
}

func (s *Service) AddTransaction(transaction *Transaction) {
	transaction.Id = s.Index
	s.Transactions = append(s.Transactions, transaction)
	s.Index++
}

func (s *Service) LastNTransactions(n int) []*Transaction {
	var resultLength int
	if len(s.Transactions) < n {
		resultLength = len(s.Transactions)
	} else {
		resultLength = n
	}

	tmp := make([]*Transaction, len(s.Transactions))
	copy(tmp, s.Transactions)

	result := make([]*Transaction, resultLength)
	for index := 0; index < resultLength; index++ {
		result[index] = tmp[len(s.Transactions)-index-1]
	}

	return result
}
