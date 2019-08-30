package bank

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/godog"
)

var testAccount *Account

func iHaveABankAccountWith(balance int) error {
	testAccount = &Account{balance: balance}
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

// Unit testing
func TestAccountDeposit(t *testing.T) {
	var account = Account{balance: 0}

	type args struct {
		amount int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"should have 10$", args{10}, 10},
		{"should have 30$", args{20}, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if account.deposit(tt.args.amount); account.balance != tt.want {
				t.Errorf("diposit() = %v, want %v", account.balance, tt.want)
			}
		})
	}
}

func TestAccountWithdraw(t *testing.T) {
	var account = Account{balance: 100}

	type args struct {
		amount int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"should have 50$", args{50}, 50},
		{"should have 0$", args{50}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if account.withdraw(tt.args.amount); account.balance != tt.want {
				t.Errorf("diposit() = %v, want %v", account.balance, tt.want)
			}
		})
	}
}
