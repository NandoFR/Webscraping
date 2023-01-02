package product

type Product struct {
	Name         string
	OriginalName string
	Category     string
	From         string
}

func NewProduct() *Product {
	return &Product{}
}
