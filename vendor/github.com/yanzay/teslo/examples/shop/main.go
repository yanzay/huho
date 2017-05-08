package main

import (
	"database/sql"
	"io"
	"strings"

	"github.com/yanzay/log"

	_ "github.com/lib/pq"
	reform "gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/yanzay/teslo"
	"github.com/yanzay/teslo/examples/shop/templates"
)

var sessions = map[string]*templates.State{}
var defaultState templates.State
var products []*templates.Product
var db *reform.DB

func main() {
	InitDB()
	defaultState = templates.State{Products: loadProducts()}
	server := teslo.NewServer()
	server.Render = func(w io.Writer) {
		templates.WritePage(w, defaultState)
	}
	server.InitSession = func(id string) {
		products := loadProducts()
		sessions[id] = &templates.State{Products: products}
	}
	server.CloseSession = func(id string) {
		delete(sessions, id)
	}
	server.Subscribe("products", ProductHandler)
	server.Start()
}

func InitDB() {
	log.Info("Creating connection to db")
	conn, err := sql.Open("postgres", "postgres://postgres@localhost:5432/shop?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	log.Info("creating reform DB")
	db = reform.NewDB(conn, postgresql.Dialect, reform.NewPrintfLogger(log.Printf))
}

func ProductHandler(s *teslo.Session, event *teslo.Event) {
	if event.Type == "click" {
		productID := strings.Split(event.ID, "-")[1]
		sessions[s.ID].Cart.Items = append(sessions[s.ID].Cart.Items, &templates.Item{Product: findProduct(productID), Quantity: 1})
		s.Respond("cart", templates.CartWidget(sessions[s.ID].Cart.Items))
	}
}

func loadProducts() []*templates.Product {
	if db == nil {
		log.Fatal("no connection to database")
	}
	db.FindAllFrom(ProductTable, "id", 1)

	products = []*templates.Product{
		{ID: "1", Name: "Oolong", Price: "2 UAH", Description: "Mi parolas la Esperanta iomete."},
		{ID: "2", Name: "Puer", Price: "3 UAH"},
	}
	return products
}

func findProduct(id string) *templates.Product {
	for _, product := range products {
		if product.ID == id {
			return product
		}
	}
	return nil
}
