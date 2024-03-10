package pkg

import (
 "fmt"
 "hw/utils"
)

type QR interface {
 Pay(amount float64) error
 CheckBalance() (float64, error)
 TransferTo(accountID int, amount float64) error
}

func DayThirteenth() {
 userWallet := &utils.OnlineWallet{Balance: 100.0}


 err := userWallet.Pay(50.0)
 if err != nil {
  fmt.Println("Ошибка оплаты:", err)
 }

 balance, _ := userWallet.CheckBalance()
 fmt.Printf("Баланс кошелька: %.2f\n", balance)

 err = userWallet.TransferTo(12345, 20.0)
 if err != nil {
  fmt.Println("Ошибка при переводе средств:", err)
 }
}
