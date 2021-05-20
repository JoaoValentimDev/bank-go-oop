package accounts

import "bank/clients"

// não temos herança em go. Temos uma composição.
type CurrentAccount struct {
	Holder                      clients.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CurrentAccount) WithdrawMoney(value float64) string {
	withdraw := value > 0 && value <= c.balance
	message := ""
	if withdraw {
		c.balance -= value
		message = "Saque efetuado com sucesso!"
	} else {
		message = "Saldo insuficiente :("
	}
	return message
}

func (c *CurrentAccount) Deposit(value float64) (status string, valueDeposit float64) {
	deposit := value > 0
	if deposit {
		status = "Dinheiro depositado com sucesso!"
		c.balance += value
		valueDeposit = c.balance
	} else {
		status = "Saudo negativo, impossivel depositar"
	}

	return status, valueDeposit
}

func (c *CurrentAccount) TransferValue(value float64, destination *CurrentAccount) bool {
	holderGreaterThanValue := c.balance > value && value > 0
	if holderGreaterThanValue {
		c.balance -= value
		destination.Deposit(value)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
