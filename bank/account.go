package bank

// Account : a simple bank account
type Account struct {
	balance int
}

func (a *Account) withdraw(amount int) {
	a.balance = a.balance - amount
}

func (a *Account) deposit(amount int) {
	a.balance = a.balance + amount
}
