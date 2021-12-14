package card

import "bank/pkg/bank/types"

// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируются.
func Total(cards []types.Card) types.Money{
	sum:=types.Money(0)
	for _, card := range cards {
		if !card.Active{
			continue
		}

		if card.Balance<=0{
			continue
		}

		sum+=card.Balance
	}
	return sum
}
