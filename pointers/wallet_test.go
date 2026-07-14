package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.50))

		assertWallet(t, wallet, Bitcoin(10.50))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 25}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertWallet(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertWallet(t, wallet, startingBalance)
	})
}

func assertWallet(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted a error but none happened")
	}

	if got != want {
		t.Errorf("want %q, but got %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("want no error, but got %q", got)
	}
}
