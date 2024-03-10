package utils

import (
	"fmt"
)


type OnlineWallet struct {
	Balance float64
}

func (w *OnlineWallet) Pay(amount float64) error {
	// Логика проведения оплаты
	if w.Balance < amount {
		return fmt.Errorf("Недостаточно средств на счете")
	}
	w.Balance -= amount
	return nil
}

func (w *OnlineWallet) CheckBalance() (float64, error) {
	// Логика проверки баланса
	return w.Balance, nil
}

func (w *OnlineWallet) TransferTo(accountID int, amount float64) error {
	// Логика перевода средств
	if w.Balance < amount {
		return fmt.Errorf("Недостаточно средств на счете для перевода")
	}
	// Логика перевода средств на другой счет
	return nil
}
