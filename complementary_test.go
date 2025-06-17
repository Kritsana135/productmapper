package productmapper_test

import (
	"testing"

	"github.com/Kritsana135/productmapper"
	"github.com/stretchr/testify/assert"
)

func TestWithComplementary(t *testing.T) {
	complementaryItems := []productmapper.ComplementaryItem{
		{
			ProductId: "WIPING-CLOTH",
			PerQty:    1,
		},
		{
			ProductId: "CLEANNER",
			PerQty:    1,
			Type:      "SUFFIX_TEXTURE",
		},
	}

	tests := []struct {
		name               string
		orders             []productmapper.CleanedOrder
		complementaryItems []productmapper.ComplementaryItem
		expected           []productmapper.CleanedOrder
	}{
		{
			name: "one order with two qty",
			orders: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
			},
			complementaryItems: complementaryItems,
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:        2,
					ProductId: "WIPING-CLOTH",
					Qty:       2,
				},
				{
					No:        3,
					ProductId: "CLEAR-CLEANNER",
					Qty:       2,
				},
			},
		},
		{
			name: "two order with two qty and different texture",
			orders: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					ProductId:  "FG0A-MATTE-IPHONE16PROMAX",
					MaterialId: "FG0A-MATTE",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "MATTE",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
			},
			complementaryItems: complementaryItems,
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:         2,
					ProductId:  "FG0A-MATTE-IPHONE16PROMAX",
					MaterialId: "FG0A-MATTE",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "MATTE",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:        3,
					ProductId: "WIPING-CLOTH",
					Qty:       4,
				},
				{
					No:        4,
					ProductId: "CLEAR-CLEANNER",
					Qty:       2,
				},
				{
					No:        5,
					ProductId: "MATTE-CLEANNER",
					Qty:       2,
				},
			},
		},
		{
			name: "two order with two qty",
			orders: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
			},
			complementaryItems: complementaryItems,
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:         2,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:        3,
					ProductId: "WIPING-CLOTH",
					Qty:       4,
				},
				{
					No:        4,
					ProductId: "CLEAR-CLEANNER",
					Qty:       4,
				},
			},
		},
		{
			name: "two order with two qty",
			orders: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
			},
			complementaryItems: []productmapper.ComplementaryItem{
				{
					ProductId: "WIPING-CLOTH",
					PerQty:    3,
				},
				{
					ProductId: "CLEANNER",
					PerQty:    2,
					Type:      "SUFFIX_TEXTURE",
				},
			},
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:         2,
					ProductId:  "FG0A-CLEAR-IPHONE16PROMAX",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  50,
					TotalPrice: 100,
				},
				{
					No:        3,
					ProductId: "WIPING-CLOTH",
					Qty:       12,
				},
				{
					No:        4,
					ProductId: "CLEAR-CLEANNER",
					Qty:       8,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			orders := productmapper.WithComplementary(test.orders, test.complementaryItems)
			assert.Equal(t, test.expected, orders)
		})
	}
}
