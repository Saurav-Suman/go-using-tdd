package interfaces

type Bank interface {
	Balance() (float64, error)
	Deposit(val float64) error
	Withdraw(val float64) error
}

type BAccount struct {
	Name          string
	TotalBalance  float64
	AccountNumber int64
}

func NewAccount(number int64) (*BAccount, error) {

	return &BAccount{
		AccountNumber: number,
	}, nil

}

func (ba *BAccount) Balance() (float64, error) {

	return ba.TotalBalance, nil

}

func (ba *BAccount) Deposit(money float64) {
	ba.TotalBalance += money
}
