package accounts

// Account struct
type Account struct {
	owner   string
	balance int
}

func Create(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0,
	}
	return &account
}

// Deposit x amount on your account
func (a Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
