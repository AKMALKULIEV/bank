package types

type Money int64

type Category string
type Status string
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)
type Payment struct{
	ID int
	Amount Money
	Category Category
	Status Status
}
type PAN string
type Currency string 

type Card struct {
	Id        int 
	PAN       PAN
	Balance   Money
	Currency  Currency
	Color     string
	Name      string
	Active    bool
	MinBalance Money

}