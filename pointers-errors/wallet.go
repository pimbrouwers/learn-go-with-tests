package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin represents a single bitcoin
type Bitcoin int

// String pretty prints Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet represents a cache of bitcoin
type Wallet struct {
	balance Bitcoin
}

// ErrInsufficentFunds represents an NSF error
var ErrInsufficentFunds = errors.New("cannot withdraw, insufficient funds")

// Balance returns the amount of bitcoin in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit adds bitcoing to the referenced wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw removes bitcoin from the referenced wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficentFunds
	}

	w.balance -= amount
	return nil
}
