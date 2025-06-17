package productmapper

import "errors"

type LineItemDetail struct {
	Qty        int
	UnitPrice  float64
	TotalPrice float64
}

var ErrInvalidUnitPrice = errors.New("invalid unit price")

func DiffusePrice(productParts []ProductParts, totalQty int, lineItemDetail LineItemDetail) ([]CleanedOrder, error) {
	if (lineItemDetail.UnitPrice * float64(lineItemDetail.Qty)) > lineItemDetail.TotalPrice {
		return nil, ErrInvalidUnitPrice
	}

	var cleanedOrders []CleanedOrder
	unitPrice := lineItemDetail.TotalPrice / float64(totalQty*lineItemDetail.Qty)

	for _, productPart := range productParts {
		qty := productPart.Qty * lineItemDetail.Qty

		cleanedOrders = append(cleanedOrders, CleanedOrder{
			ProductId:  productPart.ProductId(),
			MaterialId: productPart.MaterialId(),
			ModelId:    productPart.ModelId,
			TextureId:  productPart.TextureId,
			Qty:        qty,
			UnitPrice:  unitPrice,
			TotalPrice: unitPrice * float64(qty),
		})

	}

	return cleanedOrders, nil
}
