package bitcoin

import (
	"errors"
	"fmt"
)

// Bitcoin represents bitcoin
type Bitcoin int

// Wallet holds a bitcoin balance
type Wallet struct {
	balance Bitcoin
}

// Deposit increases the balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds returned when balance is too low
var ErrInsufficientFunds = errors.New("insuccible funds")

// Withdraw decreases the balance
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount

	return nil
}

// Balance returns the balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
