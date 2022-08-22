// Package wallet is a wallet.
package wallet

import "fmt"

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
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
