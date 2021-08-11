package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		checkBalance(t, wallet, want)
	})

	t.Run("Withdraw sucess", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertNoError(t, err)
		checkBalance(t, wallet, want)
	})

	t.Run("Withdraw insuficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(10)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(20))
		checkBalance(t, wallet, initialBalance)
		assertError(t, err, ErrInsufficientFunds.Error())
	})
}

func checkBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, expected %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
