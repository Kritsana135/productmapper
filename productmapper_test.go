package productmapper_test

import (
	"context"
	"testing"

	"github.com/Kritsana135/productmapper"
	"github.com/stretchr/testify/assert"
)

func TestCleanOrder(t *testing.T) {
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
		orders             []productmapper.InputOrder
		complementaryItems []productmapper.ComplementaryItem
		expected           []productmapper.CleanedOrder
		err                error
	}{
		{
			name: "Case 1 : Only one product",
			orders: []productmapper.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
					Qty:               2,
					UnitPrice:         50,
					TotalPrice:        100,
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
			name: "Case 2 : One product with wrong prefix",
			orders: []productmapper.InputOrder{
				{
					No:                1,
					PlatformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
					Qty:               2,
					UnitPrice:         50,
					TotalPrice:        100,
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
			name: "Case 3: One product with wrong prefix and has * symbol that indicates the quantity",
			orders: []productmapper.InputOrder{
				{
					No:                1,
					PlatformProductId: "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
					Qty:               1,
					UnitPrice:         90,
					TotalPrice:        90,
				},
			},
			complementaryItems: complementaryItems,
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-MATTE-IPHONE16PROMAX",
					MaterialId: "FG0A-MATTE",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "MATTE",
					Qty:        3,
					UnitPrice:  30,
					TotalPrice: 90,
				},
				{
					No:        2,
					ProductId: "WIPING-CLOTH",
					Qty:       3,
				},
				{
					No:        3,
					ProductId: "MATTE-CLEANNER",
					Qty:       3,
				},
			},
		},
		{
			name: "Case 4: One bundle product with wrong prefix and split by / symbol into two product",
			orders: []productmapper.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
					Qty:               1,
					UnitPrice:         80,
					TotalPrice:        80,
				},
			},

			complementaryItems: complementaryItems,
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					TextureId:  "CLEAR",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         2,
					ProductId:  "FG0A-CLEAR-OPPOA3-B",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3-B",
					TextureId:  "CLEAR",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:        3,
					ProductId: "WIPING-CLOTH",
					Qty:       2,
				},
				{
					No:        4,
					ProductId: "CLEAR-CLEANNER",
					Qty:       2,
				},
			},
		},
		{
			name: "Case 5: One bundle product with wrong prefix and split by / symbol into three product ",
			orders: []productmapper.InputOrder{
				{
					No:                1,
					PlatformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MATTE-OPPOA3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},

			complementaryItems: complementaryItems,
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					TextureId:  "CLEAR",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         2,
					ProductId:  "FG0A-CLEAR-OPPOA3-B",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3-B",
					TextureId:  "CLEAR",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:         3,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					TextureId:  "MATTE",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:        4,
					ProductId: "WIPING-CLOTH",
					Qty:       3,
				},
				{
					No:        5,
					ProductId: "CLEAR-CLEANNER",
					Qty:       2,
				},
				{
					No:        6,
					ProductId: "MATTE-CLEANNER",
					Qty:       1,
				},
			},
		},
		{
			name: "Case 6: One bundle product with wrong prefix and have / symbol and * symbol",
			orders: []productmapper.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
					Qty:               1,
					UnitPrice:         120,
					TotalPrice:        120,
				},
			},

			complementaryItems: complementaryItems,
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         2,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					TextureId:  "MATTE",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					No:        3,
					ProductId: "WIPING-CLOTH",
					Qty:       3,
				},
				{
					No:        4,
					ProductId: "CLEAR-CLEANNER",
					Qty:       2,
				},
				{
					No:        5,
					ProductId: "MATTE-CLEANNER",
					Qty:       1,
				},
			},
		},
		{
			name: "Case 7: one product and one bundle product with wrong prefix and have / symbol and * symbol",
			orders: []productmapper.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2",
					Qty:               1,
					UnitPrice:         160,
					TotalPrice:        160,
				},
				{
					No:                2,
					PlatformProductId: "FG0A-PRIVACY-IPHONE16PROMAX",
					Qty:               1,
					UnitPrice:         50,
					TotalPrice:        50,
				},
			},

			complementaryItems: complementaryItems,
			expected: []productmapper.CleanedOrder{
				{
					No:         1,
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         2,
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					TextureId:  "MATTE",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					No:         3,
					ProductId:  "FG0A-PRIVACY-IPHONE16PROMAX",
					MaterialId: "FG0A-PRIVACY",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "PRIVACY",
					Qty:        1,
					UnitPrice:  50,
					TotalPrice: 50,
				},
				{
					No:        4,
					ProductId: "WIPING-CLOTH",
					Qty:       5,
				},
				{
					No:        5,
					ProductId: "CLEAR-CLEANNER",
					Qty:       2,
				},
				{
					No:        6,
					ProductId: "MATTE-CLEANNER",
					Qty:       2,
				},
				{
					No:        7,
					ProductId: "PRIVACY-CLEANNER",
					Qty:       1,
				},
			},
		},
		// additional test case
		{
			name: "invalid platform product id: quantity symbol without number",
			orders: []productmapper.InputOrder{
				{
					No:                1,
					PlatformProductId: "--FG0A-CLEAR-OPPOA3*/FG0A-MATTE-OPPOA3*2",
					Qty:               1,
					UnitPrice:         160,
					TotalPrice:        160,
				},
			},
			complementaryItems: complementaryItems,
			err: &productmapper.ParseError{
				Message: "quantity symbol '*' found but no digits followed",
				Input:   "--FG0A-CLEAR-OPPOA3*/FG0A-MATTE-OPPOA3*2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			orders, err := productmapper.CleanOrder(context.Background(), test.orders, test.complementaryItems)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, orders)
		})
	}
}
