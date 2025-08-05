package stats

import (
	
	"github.com/Amee6/bank/v3/pkg/types"
)

func Avg(payments [] types.Payment) types.Money	{
	var avgSum types.Money
	
	for _, payment := range payments {
		if payment.Status != " FAIL" {
		avgSum += payment.Amount
	}
	}
	avgSum = avgSum / types.Money(len(payments))
	return avgSum
}

func TotalICategory(payments [] types.Payment, category types.Category) types.Money {
	
	var sum types.Money

	for _, payment := range payments {
		if payment.Category == category && payment.Status != "FAIL" {
			sum += payment.Amount
		}
	}
	return sum
}