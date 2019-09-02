package bitcoin

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Test deposit", func(t *testing.T) {
		w := Wallet{}

		w.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, w, want)
	})

	t.Run("Test withdraw", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(10)}

		err := w.Withdraw(Bitcoin(3))

		want := Bitcoin(7)

		assertBalance(t, w, want)
		assertNoError(t, err)
	})

	t.Run("Test insufficient funds", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(5)}

		err := w.Withdraw(Bitcoin(10))

		assertBalance(t, w, Bitcoin(5))

		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func assertError(t *testing.T, e error, want error) {
	t.Helper()

	if e == nil {
		t.Error("Wanted an error but didn't get one")
	}

	if e != want {
		t.Errorf("Got %q but wanted %q", e, want)
	}
}

func assertNoError(t *testing.T, e error) {
	t.Helper()

	if e != nil {
		t.Errorf("Expected no error but go one %q", e)
	}
}
