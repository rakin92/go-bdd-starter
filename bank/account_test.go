package bank

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

var testAccount *account

func iHaveABankAccountWith(balance int) error {
	testAccount = &account{balance: balance}
	return nil
}

func iDeposit(amount int) error {
	testAccount.deposit(amount)
	return nil
}

func iWithdraw(amount int) error {
	testAccount.withdraw(amount)
	return nil
}

func itShouldHaveABalanceOf(balance int) error {
	if testAccount.balance == balance {
		return nil
	}
	return fmt.Errorf("Incorrect account balance")
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have a bank account with (\d+)\$$`, iHaveABankAccountWith)
	s.Step(`^I deposit (\d+)\$$`, iDeposit)
	s.Step(`^I withdraw (\d+)\$$`, iWithdraw)
	s.Step(`^it should have a balance of (\d+)\$$`, itShouldHaveABalanceOf)

	s.BeforeScenario(func(interface{}) {
		testAccount = nil
	})
}
