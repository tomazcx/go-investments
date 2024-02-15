package accountio

import "time"

type DepositInput struct {
	ID      string
	Ammount float32
}

type DepositOutput struct {
	NewBalance float32
	Timestamp time.Time
}
