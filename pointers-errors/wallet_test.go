package wallet

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	}

	assertError := func(t *testing.T, err error, expected error) {
		t.Helper()

		if err == nil {
			t.Fatal("wanted error but didn't get one")
		}

		if err != expected {
			t.Errorf("got '%s' expected '%s", err, expected)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdaw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)

	})

	t.Run("Withdraw insufficient", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficentFunds)
	})
}
