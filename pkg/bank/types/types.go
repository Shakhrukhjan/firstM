package types

const (
	TJS currency = "TJS"
	RUB currency = "RUB"
	USD currency = "USD"
)

type currency string
type Money int
type Payment struct {
	ID     int
	Amount Money
}
