package productmapper_test

import (
	"testing"

	"github.com/Kritsana135/productmapper"
	"github.com/stretchr/testify/assert"
)

func TestDiffusePrice(t *testing.T) {
	tests := []struct {
		name             string
		productParts     []productmapper.ProductParts
		lineItemDetail   productmapper.LineItemDetail
		totalQty         int
		expectedProducts []productmapper.CleanedOrder
		expectedError    error
	}{
		{
			name: "one product parts",
			productParts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        1,
				},
			},
			totalQty: 1,
			lineItemDetail: productmapper.LineItemDetail{
				Qty:        2,
				UnitPrice:  50,
				TotalPrice: 100,
			},
			expectedProducts: []productmapper.CleanedOrder{
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
		},
		{
			name: "balance unit price",
			productParts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "MATTE",
					ModelId:    "IPHONE16PROMAX",
					Qty:        3,
				},
			},
			totalQty: 3,
			lineItemDetail: productmapper.LineItemDetail{
				Qty:        1,
				UnitPrice:  90,
				TotalPrice: 90,
			},
			expectedProducts: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-MATTE-IPHONE16PROMAX",
					MaterialId: "FG0A-MATTE",
					ModelId:    "IPHONE16PROMAX",
					TextureId:  "MATTE",
					Qty:        3,
					UnitPrice:  30,
					TotalPrice: 90,
				},
			},
		},
		{
			name: "1: balance unit price and total price for two product parts",
			productParts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3",
					Qty:        1,
				},
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3-B",
					Qty:        1,
				},
			},
			totalQty: 2,
			lineItemDetail: productmapper.LineItemDetail{
				Qty:        1,
				UnitPrice:  80,
				TotalPrice: 80,
			},
			expectedProducts: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					TextureId:  "CLEAR",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					ProductId:  "FG0A-CLEAR-OPPOA3-B",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3-B",
					TextureId:  "CLEAR",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
			},
		},
		{
			name: "2: balance unit price and total price for two product parts",
			productParts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3",
					Qty:        2,
				},
				{
					FilmTypeId: "FG0A",
					TextureId:  "MATTE",
					ModelId:    "OPPOA3",
					Qty:        1,
				},
			},
			totalQty: 3,
			lineItemDetail: productmapper.LineItemDetail{
				Qty:        1,
				UnitPrice:  120,
				TotalPrice: 120,
			},
			expectedProducts: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					TextureId:  "MATTE",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
			},
		},
		{
			name: "3: balance unit price and total price for two product parts",
			productParts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3",
					Qty:        2,
				},
				{
					FilmTypeId: "FG0A",
					TextureId:  "MATTE",
					ModelId:    "OPPOA3",
					Qty:        2,
				},
			},
			totalQty: 4,
			lineItemDetail: productmapper.LineItemDetail{
				Qty:        1,
				UnitPrice:  160,
				TotalPrice: 160,
			},
			expectedProducts: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					TextureId:  "CLEAR",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
				{
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					TextureId:  "MATTE",
					Qty:        2,
					UnitPrice:  40,
					TotalPrice: 80,
				},
			},
		},
		{
			name: "balance unit price and total price for three product parts",
			productParts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3",
					Qty:        1,
				},
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3-B",
					Qty:        1,
				},
				{
					FilmTypeId: "FG0A",
					TextureId:  "MATTE",
					ModelId:    "OPPOA3",
					Qty:        1,
				},
			},
			totalQty: 3,
			lineItemDetail: productmapper.LineItemDetail{
				Qty:        1,
				UnitPrice:  120,
				TotalPrice: 120,
			},
			expectedProducts: []productmapper.CleanedOrder{
				{
					ProductId:  "FG0A-CLEAR-OPPOA3",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3",
					TextureId:  "CLEAR",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					ProductId:  "FG0A-CLEAR-OPPOA3-B",
					MaterialId: "FG0A-CLEAR",
					ModelId:    "OPPOA3-B",
					TextureId:  "CLEAR",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
				{
					ProductId:  "FG0A-MATTE-OPPOA3",
					MaterialId: "FG0A-MATTE",
					ModelId:    "OPPOA3",
					TextureId:  "MATTE",
					Qty:        1,
					UnitPrice:  40,
					TotalPrice: 40,
				},
			},
		},
		{
			name: "if unit price > total price in line item detail should return error",
			productParts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        1,
				},
			},
			totalQty: 1,
			lineItemDetail: productmapper.LineItemDetail{
				Qty:        2,
				UnitPrice:  150,
				TotalPrice: 100,
			},
			expectedError: productmapper.ErrInvalidUnitPrice,
		},
		{
			name: "if unit price * qty > total price in line item detail should return error",
			productParts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        1,
				},
			},
			totalQty: 1,
			lineItemDetail: productmapper.LineItemDetail{
				Qty:        2,
				UnitPrice:  51,
				TotalPrice: 100,
			},
			expectedError: productmapper.ErrInvalidUnitPrice,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			products, err := productmapper.DiffusePrice(test.productParts, test.totalQty, test.lineItemDetail)

			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedProducts, products)

		})
	}
}
