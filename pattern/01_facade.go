package main

import (
	"fmt"
	"log"
)

// 1 подсистема
type account struct {
	name string
}

func NewAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

// 2 подсистема
type securityCode struct {
	code int
}

func NewSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) CheckCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("SecurityCode Verified")
	return nil
}

// 3 подсистема
type wallet struct {
	balance int
}

func NewWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Println("Wallet balance is sufficient")
	w.balance = w.balance - amount
	return nil
}

// 4 подсистема
type ledger struct {
}

func NewLedger() *ledger {
	return &ledger{}
}

func (l *ledger) MakeEntry(accountID string, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
}

// 5 подсистема
type notification struct {
}

func NewNotification() *notification {
	return &notification{}
}

func (n *notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

// Фасад.
type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func NewWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletF := &walletFacade{
		account:      NewAccount(accountID),
		securityCode: NewSecurityCode(code),
		wallet:       NewWallet(),
		notification: NewNotification(),
		ledger:       NewLedger(),
	}
	fmt.Println("Account created")
	return walletF
}

func (w *walletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.SendWalletCreditNotification()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.SendWalletDebitNotification()
	w.ledger.MakeEntry(accountID, "debit", amount)
	return nil
}

func main() {
	fmt.Println()
	walletF := NewWalletFacade("abc", 1234)
	fmt.Println()
	err := walletF.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletF.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
