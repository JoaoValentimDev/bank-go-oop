package accounts

import "bank/clients"

type SavingsAccount struct {
	Holder                                 clients.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (s *SavingsAccount) GetBalance() float64 {
	return s.balance
}

func (s *SavingsAccount) WithdrawMoney(value float64) string {
	withdraw := value > 0 && value <= s.balance
	message := ""
	if withdraw {
		s.balance -= value
		message = "Saque efetuado com sucesso!"
	} else {
		message = "Saldo insuficiente :("
	}
	return message
}

func (s *SavingsAccount) Deposit(value float64) (status string, valueDeposit float64) {
	deposit := value > 0
	if deposit {
		status = "Dinheiro depositado com sucesso!"
		s.balance += value
		valueDeposit = s.balance
	} else {
		status = "Saudo negativo, impossivel depositar"
	}

	return status, valueDeposit
}

func (s *SavingsAccount) TransferValue(value float64, destination *CurrentAccount) bool {
	holderGreaterThanValue := s.balance > value && value > 0
	if holderGreaterThanValue {
		s.balance -= value
		destination.Deposit(value)
		return true
	} else {
		return false
	}
}
