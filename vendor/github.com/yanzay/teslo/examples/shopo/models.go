//go:generate reform
package main

//reform:products
type Product struct {
	ID          int32  `reform:"id,pk"`
	Name        string `reform:"name"`
	Description string `reform:"description"`
	Price       int    `reform:"price"`
}

//reform:orders
type Order struct {
	ID      int32  `reform:"id,pk"`
	Address string `reform:"address"`
	Phone   string `reform:"phone"`
}

//reform:items
type Item struct {
	ID        int32 `reform:"id,pk"`
	OrderID   int32 `reform:"order_id"`
	ProductID int32 `reform:"product_id"`
	Quantity  int   `reform:"quantity"`
}
