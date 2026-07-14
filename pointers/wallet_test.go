package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertWallet := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		fmt.Printf("the address of the wallet in test is: %p \n", &wallet.balance)

		got := wallet.Balance()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.50))

		assertWallet(t, wallet, Bitcoin(10.50))

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 25}
		wallet.Withdraw(Bitcoin(10))

		assertWallet(t, wallet, Bitcoin(15))
	})
}
