// Package wallet is a wallet.
package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet_Balance(t *testing.T) {
	t.Parallel()
	type fields struct {
		balance Bitcoin
	}
	tests := []struct {
		name   string
		fields fields
		want   Bitcoin
	}{
		{"balance", fields{10}, 10},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := Wallet{
				balance: test.fields.balance,
			}
			assert.Equal(t, test.want, w.Balance())
		})
	}
}

func TestWallet_Deposit(t *testing.T) {
	t.Parallel()
	type fields struct {
		balance Bitcoin
	}
	type args struct {
		amount Bitcoin
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bitcoin
	}{
		{"+deposit", fields{10}, args{10}, 20},
		{"-deposit", fields{10}, args{-10}, 0},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := Wallet{
				balance: test.fields.balance,
			}
			w.Deposit(test.args.amount)
			assert.Equal(t, test.want, w.Balance())
		})
	}
}

func TestBitcoin_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		b    Bitcoin
		want string
	}{
		{"btc", Bitcoin(5), "5 BTC"},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.want, test.b.String())
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	t.Parallel()
	type fields struct {
		balance Bitcoin
	}
	type args struct {
		amount Bitcoin
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bitcoin
	}{
		{"withdraw", fields{5}, args{3}, 2},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			w := &Wallet{
				balance: test.fields.balance,
			}
			w.Withdraw(test.args.amount)
			assert.Equal(t, test.want, w.Balance())
		})
	}
}
