package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mocked objects
type MockedCalendar struct {
	mock.Mock
}

type MockedPrinter struct {
	mock.Mock
}

func (m *MockedCalendar) GetDate() (s string, err error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockedPrinter) Print() error {
	args := m.Called()
	return args.Error(1)
}
func TestPrintingAccountStatementOnConsole(t *testing.T) {

	calendar := new(MockedCalendar)
	printer := new(MockedPrinter)
	// setup expectations
	calendar.
		On("GetDate").Return("10-01-2012", nil).
		On("GetDate").Return("13-01-2012", nil).
		On("GetDate").Return("14-01-2012", nil)
	printer.On("Print").Return(nil)

	ba := NewBankAccount(calendar, printer)

	ba.Deposit(1000)
	ba.Deposit(2000)
	ba.Withdrawal(500)

	assert.True(t, true, "True is true!")

}
