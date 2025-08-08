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
    sumByCategory := map[types.Category]types.Money{}
    countByCategory := map[types.Category]int{}

    for _, payment := range payments {
        sumByCategory[payment.Category] += payment.Amount
        countByCategory[payment.Category]++
    }

    avgByCategory := map[types.Category]types.Money{}
    for category, sum := range sumByCategory {
        avgByCategory[category] = sum / types.Money(countByCategory[category])
    }

    return avgByCategory
}

// PeriodsDynamic сравнивает расходы по категориям за два периода
func PeriodsDynamic (first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	
	third := map[types.Category]types.Money{}

	for category, amountFirst := range first {
        amountSecond := second[category]
        third[category] = amountSecond - amountFirst
    }

    for category, amountSecond := range second {
        if _, exists := first[category]; !exists {
            third[category] = amountSecond
        }
    }

    return third
	
}