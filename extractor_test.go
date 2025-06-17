package productmapper_test

import (
	"testing"

	"github.com/Kritsana135/productmapper"
	"github.com/stretchr/testify/assert"
)

func TestExtractPlatformId(t *testing.T) {
	tests := []struct {
		name              string
		platformProductId string
		expectedProducts  []productmapper.ProductParts
		totalQty          int
		err               error
	}{
		// problem test case
		{
			name:              "one product id with correct format",
			platformProductId: "FG0A-CLEAR-IPHONE16PROMAX",
			expectedProducts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        1,
				},
			},
			totalQty: 1,
		},
		{
			name:              "one product with wrong prefix",
			platformProductId: "x2-3&FG0A-CLEAR-IPHONE16PROMAX",
			expectedProducts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "IPHONE16PROMAX",
					Qty:        1,
				},
			},
			totalQty: 1,
		},
		{
			name:              "one product with wrong prefix and quantity symbol",
			platformProductId: "x2-3&FG0A-MATTE-IPHONE16PROMAX*3",
			expectedProducts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "MATE",
					ModelId:    "IPHONE16PROMAX",
					Qty:        3,
				},
			},
			totalQty: 3,
		},
		{
			name:              "one bundle product with wrong prefix and split by / symbol into two product",
			platformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B",
			expectedProducts: []productmapper.ProductParts{
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
		},
		{
			name:              "one bundle product with wrong prefix and split by / symbol into three product",
			platformProductId: "FG0A-CLEAR-OPPOA3/%20xFG0A-CLEAR-OPPOA3-B/FG0A-MATTE-OPPOA3",
			expectedProducts: []productmapper.ProductParts{
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
		},
		{
			name:              "one bundle product with wrong prefix and have / symbol and * symbol",
			platformProductId: "--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3",
			expectedProducts: []productmapper.ProductParts{
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
		},
		// additional test case
		{
			name:              "one product that has product type",
			platformProductId: "FG0A-CLEAR-OPPOA3-B",
			expectedProducts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3-B",
					Qty:        1,
				},
			},
			totalQty: 1,
		},
		{
			name:              "one product that has product type and quantity symbol",
			platformProductId: "FG0A-CLEAR-OPPOA3-B*730",
			expectedProducts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3-B",
					Qty:        730,
				},
			},
			totalQty: 730,
		},
		{
			name:              "two product and split by /",
			platformProductId: "FG0A-CLEAR-OPPOA3-B*710/FI2A-MATE-NOKIA3310",
			expectedProducts: []productmapper.ProductParts{
				{
					FilmTypeId: "FG0A",
					TextureId:  "CLEAR",
					ModelId:    "OPPOA3-B",
					Qty:        710,
				},
				{
					FilmTypeId: "FI2A",
					TextureId:  "MATE",
					ModelId:    "NOKIA3310",
					Qty:        1,
				},
			},
			totalQty: 711,
		},
		{
			name:              "invalid texture id format",
			platformProductId: "FG0A-CLEAR*2-OPPOA3-B",
			expectedProducts:  []productmapper.ProductParts{},
			err: &productmapper.ParseError{
				Message: "invalid texture id format",
				Input:   "FG0A-CLEAR*2-OPPOA3-B",
				Index:   10,
			},
			totalQty: 0,
		},
		{
			name:              "missing model section",
			platformProductId: "FG0A-CLEAR-",
			expectedProducts:  []productmapper.ProductParts{},
			err: &productmapper.ParseError{
				Message: "invalid format",
				Input:   "FG0A-CLEAR-",
			},
			totalQty: 0,
		},
		{
			name:              "one bundle product seperate by / first section is invalid format",
			platformProductId: "FG0A-CLEAR-/FI2A-MATE-NOKIA3310",
			expectedProducts:  []productmapper.ProductParts{},
			err: &productmapper.ParseError{
				Message: "invalid format",
				Input:   "FG0A-CLEAR-/FI2A-MATE-NOKIA3310",
			},
			totalQty: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			products, totalQty, err := productmapper.ExtractPlatformId(tc.platformProductId)

			assert.Equal(t, tc.err, err)
			assert.Len(t, products, len(tc.expectedProducts), "Expected number of products to match")
			assert.Equal(t, tc.totalQty, totalQty, "Expected total quantity to match")

			for i, expected := range tc.expectedProducts {
				assert.Equal(t, expected.FilmTypeId, products[i].FilmTypeId, "MaterialId mismatch for product %d", i)
				assert.Equal(t, expected.ModelId, products[i].ModelId, "ModelId mismatch for product %d", i)
				assert.Equal(t, expected.Qty, products[i].Qty, "Quantity mismatch for product %d", i)
			}
		})
	}
}
