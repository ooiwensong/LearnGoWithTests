package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(Bitcoin(10))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(0))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	// t.Fatal will stop the test if error == nil, this prevents the code from
	// continuing and getting a nil pointer error.
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	// .Error() converts errors to strings
	// if got.Error() != want {
	// 	t.Errorf("got %q, want %q", got, want)
	// }

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
