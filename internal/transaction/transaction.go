package transaction

type Entry struct {
	AccountID string
	Amount int64
}

type Transaction struct {
	ID string
	Entries []Entry
}

func (t Transaction)IsBalanced() bool{
	var total int64
	for _, entry :=range(t.Entries){
		total+=entry.Amount
	}
	return total==0;
}