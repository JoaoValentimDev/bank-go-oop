package main

import (
	"bank/accounts"
	"bank/clients"
	"fmt"
)

type verifyAccount interface {
	WithdrawMoney(value float64) string
}

func PaySlip(conta verifyAccount, valueSlip float64) {
	conta.WithdrawMoney(valueSlip)
}

func main() {
	// para passar um valor de uma conta para a outra, usa-se & para dizer o endereço o endereço
	holder := clients.Holder{
		Name:       "João",
		CPF:        "222.222.222-22",
		Profession: "Developer",
	}

	currentAccount := accounts.CurrentAccount{
		Holder:        holder,
		AgencyNumber:  2,
		AccountNumber: 2,
	}

	currentAccount.Deposit(2000)

	PaySlip(&currentAccount, 200)

	fmt.Println(currentAccount)

	savingsAccount := accounts.SavingsAccount{
		Holder:        holder,
		AgencyNumber:  222,
		AccountNumber: 241,
		Operation:     2,
	}

	savingsAccount.Deposit(1000)

	// Pagamos um boleto apartir deste método, com & referenciamos a struct,
	// que tera o mesmo método que foi moldado na interface
	PaySlip(&savingsAccount, 500)

	fmt.Println(savingsAccount.GetBalance())
}
