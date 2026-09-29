package ledger

import {
	"fmt"
	"github.com/choijiwonsoc/ledger/internal/account"
	"github.com/choijiwonsoc/ledger/internal/transaction"
}

type Ledger struct{
	Accounts map[string]*account.Account
	Transactions map[string]*transaction.Transaction
}

func New() *Ledger{
	return &Ledger{
		Accounts:make(map[string]*account.Account),
		Transactions:make(map[string]*transaction.Transaction)
	}
}
\
func(l *Ledger) CreateAccount(id string, name string) error{
	if _, exists:=l.Accounts[id].exists{
		return fmt.Errorf("Account %s alr exists", id)
	}
	l.Accounts[id] = account.New(id, name)
	return nil
}