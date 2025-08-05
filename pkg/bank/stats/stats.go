package stats

import "github.com/Amee6/bank/pkg/types"

func Avg(payments [] types.Payment) types.Money	{
	var avgSum types.Money

	for _, payment := range payments {
		avgSum += payment.Amount
	}
	avgSum = avgSum / types.Money(len(payments))
	return avgSum
}

func TotalICategory(payments [] types.Payment, category types.Category) types.Money {
	
	var sum types.Money

	for _, payment := range payments {
		if payment.Category == category{
			sum += payment.Amount
		}
	}
	return sum
}