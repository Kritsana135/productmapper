package productmapper

import "github.com/elliotchance/orderedmap/v3"

type ComplementaryItem struct {
	ProductId string
	PerQty    int
	// for logical complementary item
	Type string // SUFFIX_TEXTURE
}

func WithComplementary(orders []CleanedOrder, complementaryItems []ComplementaryItem) []CleanedOrder {
	newOrders := []CleanedOrder{}
	omapComplementary := orderedmap.NewOrderedMap[string, int]()

	orderNo := 1
	for _, order := range orders {
		order.No = orderNo
		orderNo++
		newOrders = append(newOrders, order)

		for _, complementaryItem := range complementaryItems {
			switch complementaryItem.Type { // implement more type later
			case "SUFFIX_TEXTURE":
				key := order.TextureId + "-" + complementaryItem.ProductId
				currentQty, _ := omapComplementary.Get(key)
				omapComplementary.Set(key, currentQty+order.Qty*complementaryItem.PerQty)
			default:
				key := complementaryItem.ProductId
				currentQty, _ := omapComplementary.Get(key)
				omapComplementary.Set(key, currentQty+order.Qty*complementaryItem.PerQty)
			}
		}
	}

	for productId, qty := range omapComplementary.AllFromFront() {
		newOrders = append(newOrders, CleanedOrder{
			No:        orderNo,
			ProductId: productId,
			Qty:       qty,
		})
		orderNo++
	}

	return newOrders
}
