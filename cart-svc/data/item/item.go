package item

import (
	d "github.com/barshat7/foodiecart/data"
)

const (
	createItemQuery string = "INSERT cart_items SET cart_id = ?, item_id = ?, item_name = ?, item_price = ?, quantity = ?"
	findByCartIDQuery string = "SELECT id, cart_id, item_id, item_name, item_price,quantity FROM cart_items where cart_id = ?"
)

// Item Items in a cart
type Item struct {
	ID int64
	CartID int64
	ItemID int64
	ItemName string
	ItemPrice float64
	Quantity int64
}

// New Creats new Item
func New (itemID int64, itemName string, itemPrice float64, quantity int64) *Item {
	return &Item{ItemID: itemID, ItemName: itemName, ItemPrice: itemPrice, Quantity: quantity}
}

// Save Saves the Item
func (item *Item) Save() (int64, error) {
	stmt, err := d.Db.Prepare(createItemQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(item.CartID, item.ItemID, item.ItemName, item.ItemPrice, item.Quantity)
	lastInserted, err := res.LastInsertId()
	return lastInserted, err
}

func FindByCartID(cartID int64) (items [] Item) {
	rows, err := d.Db.Query(findByCartIDQuery, cartID)
	if (err != nil) {
		panic(err)
	}
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.CartID, &item.ItemID, &item.ItemName, &item.ItemPrice, &item.Quantity)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return items
}