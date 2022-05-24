package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {

	payments := []types.Payment{
		{
			ID:     1,
			Amount: 17_000_00,
		},
		{
			ID:     2,
			Amount: 14_000_00,
		},
		{
			ID:     3,
			Amount: 15_000_00,
		},
		{
			ID:     4,
			Amount: 25_000_00,
		},
		{
			ID:     5,
			Amount: 24_000_00,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)

	// Output: 4
}
