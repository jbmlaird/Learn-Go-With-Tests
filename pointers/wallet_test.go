package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		checkBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		checkBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}

func checkBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	if wallet.Balance() != want {
		t.Errorf("got %s want %s", wallet.Balance(), want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		// Stop the test here if not existing as nil pointer will be thrown next assertion
		t.Fatal("Wanted error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Didn't want an error but got one")
	}
}
