package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance:1_000_00,
			Active:true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance:1_000_00,
			Active:true,
		},
		{
			Balance:2_000_00,
			Active:true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance:1_000_00,
			Active:false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance:-1_000_00,
			Active:true,
		},
	}))

	// Output:
	// 100000
	// 300000
	// 0
	// 0
}

func ExamplePaymentSources() {
	sources:=PaymentSources([]types.Card{
		{
			Balance:1_000_00,
			Active:true,
			PAN: "5058 xxxx xxxx 2020",
		},
		{
			Balance:1_000_00,
			Active:false,
			PAN: "5058 xxxx xxxx 2021",
		},
		{
			Balance:0,
			Active:true,
			PAN: "5058 xxxx xxxx 2022",
		},
		{
			Balance:-1_000_00,
			Active:true,
			PAN: "5058 xxxx xxxx 2023",
		},
		{
			Balance:1_000_00,
			Active:true,
			PAN: "5058 xxxx xxxx 2024",
		},
	})

	for _, source := range sources {
		fmt.Println(source.Number)
	}

	// Output: 
	// 5058 xxxx xxxx 2020
	// 5058 xxxx xxxx 2024
}