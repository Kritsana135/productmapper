package productmapper

import "context"

type InputOrder struct {
	No                int
	PlatformProductId string
	Qty               int
	UnitPrice         float64
	TotalPrice        float64
}

type CleanedOrder struct {
	No         int
	TextureId  string
	ProductId  string
	MaterialId string
	ModelId    string
	Qty        int
	UnitPrice  float64
	TotalPrice float64
}

func CleanOrder(ctx context.Context, orders []InputOrder, complementaryItems []ComplementaryItem) ([]CleanedOrder, error) {
	var cleanedOrders []CleanedOrder

	for _, order := range orders {
		productParts, totalQty, err := ExtractPlatformId(order.PlatformProductId)
		if err != nil {
			return nil, err
		}

		diffusedOrders, err := DiffusePrice(productParts, totalQty, LineItemDetail{
			Qty:        order.Qty,
			UnitPrice:  order.UnitPrice,
			TotalPrice: order.TotalPrice,
		})
		if err != nil {
			return nil, err
		}

		cleanedOrders = append(cleanedOrders, diffusedOrders...)
	}

	return WithComplementary(cleanedOrders, complementaryItems), nil
}
