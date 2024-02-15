package accountio

import "time"

type WithdrawInput struct {
	ID      string
	Ammount float32
}

type WithdrawOutput struct {
	NewBalance float32
	Timestamp time.Time
}
