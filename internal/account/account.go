package account

type Account struct{
	ID string
	Name string
	Balance int64
}

func New(id string, name string) *Account{
	return &Account{
		ID: id,
		Name: name,
		Balance: 0,
	}
}