package payment

import (
	"bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment {

	maxPay := payments[0]

	for i := 1; i < len(payments); i++ {
		if maxPay.Amount < payments[i].Amount {

			maxPay = payments[i]

		}
		payments[0] = maxPay
	}
	return payments[0]

}
