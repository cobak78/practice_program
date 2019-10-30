package main

import "log"

type Scheduler interface {
	GetDate() (string, error)
}

type OutputPrinter interface {
	Print() error
}

type bankAccount struct {
	calendar Scheduler
	printer  OutputPrinter
}

func NewBankAccount(s Scheduler, p OutputPrinter) bankAccount {
	return bankAccount{
		calendar: s,
		printer:  p,
	}
}
func (b *bankAccount) Deposit(amount int) error {
	d, _ := b.calendar.GetDate()
	log.Println(d)
	return nil
}
func (b *bankAccount) Withdrawal(amount int) error {
	d, _ := b.calendar.GetDate()
	log.Println(d)
	return nil
}
func (b *bankAccount) Print() error {
	return b.printer.Print()
}
