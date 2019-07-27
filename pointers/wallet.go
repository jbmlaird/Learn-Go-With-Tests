package pointers

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type String interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
