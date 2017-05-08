package main

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type productTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *productTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("products").
func (v *productTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *productTableType) Columns() []string {
	return []string{"id", "name", "description", "price"}
}

// NewStruct makes a new struct for that view or table.
func (v *productTableType) NewStruct() reform.Struct {
	return new(Product)
}

// NewRecord makes a new record for that table.
func (v *productTableType) NewRecord() reform.Record {
	return new(Product)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *productTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ProductTable represents products view or table in SQL database.
var ProductTable = &productTableType{
	s: parse.StructInfo{Type: "Product", SQLSchema: "", SQLName: "products", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "Description", PKType: "", Column: "description"}, {Name: "Price", PKType: "", Column: "price"}}, PKFieldIndex: 0},
	z: new(Product).Values(),
}

// String returns a string representation of this struct or record.
func (s Product) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "Description: " + reform.Inspect(s.Description, true)
	res[3] = "Price: " + reform.Inspect(s.Price, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Product) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Name,
		s.Description,
		s.Price,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Product) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Name,
		&s.Description,
		&s.Price,
	}
}

// View returns View object for that struct.
func (s *Product) View() reform.View {
	return ProductTable
}

// Table returns Table object for that record.
func (s *Product) Table() reform.Table {
	return ProductTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Product) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Product) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Product) HasPK() bool {
	return s.ID != ProductTable.z[ProductTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Product) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = ProductTable
	_ reform.Struct = new(Product)
	_ reform.Table  = ProductTable
	_ reform.Record = new(Product)
	_ fmt.Stringer  = new(Product)
)

type orderTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *orderTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("orders").
func (v *orderTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *orderTableType) Columns() []string {
	return []string{"id", "address", "phone"}
}

// NewStruct makes a new struct for that view or table.
func (v *orderTableType) NewStruct() reform.Struct {
	return new(Order)
}

// NewRecord makes a new record for that table.
func (v *orderTableType) NewRecord() reform.Record {
	return new(Order)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *orderTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// OrderTable represents orders view or table in SQL database.
var OrderTable = &orderTableType{
	s: parse.StructInfo{Type: "Order", SQLSchema: "", SQLName: "orders", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "Address", PKType: "", Column: "address"}, {Name: "Phone", PKType: "", Column: "phone"}}, PKFieldIndex: 0},
	z: new(Order).Values(),
}

// String returns a string representation of this struct or record.
func (s Order) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Address: " + reform.Inspect(s.Address, true)
	res[2] = "Phone: " + reform.Inspect(s.Phone, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Order) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Address,
		s.Phone,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Order) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Address,
		&s.Phone,
	}
}

// View returns View object for that struct.
func (s *Order) View() reform.View {
	return OrderTable
}

// Table returns Table object for that record.
func (s *Order) Table() reform.Table {
	return OrderTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Order) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Order) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Order) HasPK() bool {
	return s.ID != OrderTable.z[OrderTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Order) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = OrderTable
	_ reform.Struct = new(Order)
	_ reform.Table  = OrderTable
	_ reform.Record = new(Order)
	_ fmt.Stringer  = new(Order)
)

type itemTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *itemTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("items").
func (v *itemTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *itemTableType) Columns() []string {
	return []string{"id", "order_id", "product_id", "quantity"}
}

// NewStruct makes a new struct for that view or table.
func (v *itemTableType) NewStruct() reform.Struct {
	return new(Item)
}

// NewRecord makes a new record for that table.
func (v *itemTableType) NewRecord() reform.Record {
	return new(Item)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *itemTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// ItemTable represents items view or table in SQL database.
var ItemTable = &itemTableType{
	s: parse.StructInfo{Type: "Item", SQLSchema: "", SQLName: "items", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "OrderID", PKType: "", Column: "order_id"}, {Name: "ProductID", PKType: "", Column: "product_id"}, {Name: "Quantity", PKType: "", Column: "quantity"}}, PKFieldIndex: 0},
	z: new(Item).Values(),
}

// String returns a string representation of this struct or record.
func (s Item) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "OrderID: " + reform.Inspect(s.OrderID, true)
	res[2] = "ProductID: " + reform.Inspect(s.ProductID, true)
	res[3] = "Quantity: " + reform.Inspect(s.Quantity, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Item) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.OrderID,
		s.ProductID,
		s.Quantity,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Item) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.OrderID,
		&s.ProductID,
		&s.Quantity,
	}
}

// View returns View object for that struct.
func (s *Item) View() reform.View {
	return ItemTable
}

// Table returns Table object for that record.
func (s *Item) Table() reform.Table {
	return ItemTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Item) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Item) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Item) HasPK() bool {
	return s.ID != ItemTable.z[ItemTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Item) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = ItemTable
	_ reform.Struct = new(Item)
	_ reform.Table  = ItemTable
	_ reform.Record = new(Item)
	_ fmt.Stringer  = new(Item)
)

func init() {
	parse.AssertUpToDate(&ProductTable.s, new(Product))
	parse.AssertUpToDate(&OrderTable.s, new(Order))
	parse.AssertUpToDate(&ItemTable.s, new(Item))
}
