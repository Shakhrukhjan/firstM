package main

import (
	"bank/pkg/bank/payment"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 15_000_00,
		},
		{
			ID:     2,
			Amount: 20_000_00,
		},
		{
			ID:     3,
			Amount: 30_000_00,
		},
		{
			ID:     4,
			Amount: 25_000_00,
		},
		{
			ID:     5,
			Amount: 17_000_00,
		},
	}
	max := payment.Max(payments)
	fmt.Println(max.ID)
}
