package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10.50))

		got := wallet.Balance()
		want := Bitcoin(10.50)

		fmt.Printf("the address of the wallet in test is: %p \n", &wallet.balance)

		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})
}
