package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw")

// Create account
func Create(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0,
	}
	return &account
}

// Deposit x amount on your account
// NOTE: Go 의 method 는 receiver 를 복사하여 사용한다. 이로 인해 모든 객체의 불변성을 보장할 수 있다.
func (a Account) Deposit(amount int) *Account {
	a.balance += amount
	return &a
}

// Owner of your account
func (a Account) Owner() string {
	return a.owner
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
// NOTE: if you want mutable object, you can use pointer "*"
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
