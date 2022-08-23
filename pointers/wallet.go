// Package wallet is a wallet.
package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is cool.
type Bitcoin int

// String representation of Bitcoin.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is cool.
type Wallet struct {
	balance Bitcoin
}

// Balance of wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit Bitcoins wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw Bitcoins wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientBalance
	}
	w.balance -= amount

	return nil
}

// ErrInsufficientBalance when there be no money.
var ErrInsufficientBalance = errors.New("insufficient balance")
