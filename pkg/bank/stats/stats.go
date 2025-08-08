package stats

import (
	"github.com/Amee6/bank/v3/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var avgSum types.Money

	for _, payment := range payments {
		if payment.Status != " FAIL" {
			avgSum += payment.Amount
		}
	}
	avgSum = avgSum / types.Money(len(payments))
	return avgSum
}

func TotalICategory(payments []types.Payment, category types.Category) types.Money {

	var sum types.Money

	for _, payment := range payments {
		if payment.Category == category && payment.Status != "FAIL" {
			sum += payment.Amount
		}
	}
	return sum
}

// FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {

	var filtered []types.Payment

	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

// TotalInCategory возвращает сумму платежей по каждой категории
func TotalInCategories(payments [] types.Payment) map [types.Category]types.Money {
	
	categories := map [types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	return  categories	
}

// CategoriesAvg считает среднюю сумму платежа по кажой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
    sumByCategory := make(map[types.Category]types.Money)
    countByCategory := make(map[types.Category]int)

    for _, payment := range payments {
        sumByCategory[payment.Category] += payment.Amount
        countByCategory[payment.Category]++
    }

    avgByCategory := make(map[types.Category]types.Money)
    for category, sum := range sumByCategory {
        avgByCategory[category] = sum / types.Money(countByCategory[category])
    }

    return avgByCategory
}