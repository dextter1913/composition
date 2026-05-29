package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

type Items []Item

func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}

	return total
}
