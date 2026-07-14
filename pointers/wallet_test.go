package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertWallet := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	}

	assertError := func(t *testing.T, err error) {
		t.Helper()

		if err == nil {
			t.Error("Wanted a error but got nil")
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

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err)
		assertWallet(t, wallet, startingBalance)

	})
}
