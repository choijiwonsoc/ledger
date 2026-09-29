package transaction

import "testing"

func TestBalancedTransaction(t *testing.T){
	tx:=Transaction{
		ID:"tx_001",
		Entries:[]Entry{
			{
				AccountID:"acct_alice",
				Amount: 500,
			},
			{
				AccountID:"acct_bob",
				Amount: -500,
			},
		},
	}
	if !tx.IsBalanced(){
		t.Error("Expected txn to be balanced")
	}
}

func TestUnbalancedTransaction(t *testing.T){
	tx:=Transaction{
		ID:"tx_002",
		Entries:[]Entry{
			{
				AccountID:"acct_alice",
				Amount: 400,
			},
			{
				AccountID:"acct_bob",
				Amount: -500,
			},
		},
	}
	if !tx.IsBalanced(){
		t.Error("Expected txn to be balanced")
	}
}