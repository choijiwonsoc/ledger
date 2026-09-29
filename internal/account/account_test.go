package account

import "testing"

func TestNewAccount(t *testing.T){
	acc:=New("acct_alice", "Alice")

	if acc.ID != "acct_alice"{
		t.Errorf("unexpected ID %s", acc.ID);
	}
	if acc.Name != "Alice"{
		t.Errorf("unexpected name %s", acc.Name);
	}
	if acc.Balance != 0{
		t.Errorf("unexpected balance %d", acc.Balance);
	}
}