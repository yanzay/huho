package templates

type State struct {
	Products []*Product
	Cart     Cart
}

type Product struct {
	ID          string
	Name        string
	Price       string
	Description string
}

type Cart struct {
	Items []*Item
}

type Item struct {
	Product  *Product
	Quantity int
}
