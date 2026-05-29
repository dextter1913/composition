package invoice

import (
	"github.com/dextter1913/composition/pkg/customer"
	"github.com/dextter1913/composition/pkg/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items // con los corchetes se indica relacion de uno a muchos
}

func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// Set Total is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
