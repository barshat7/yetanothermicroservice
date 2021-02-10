package cart

import (
	"time"
	d "github.com/barshat7/foodiecart/data"
	i "github.com/barshat7/foodiecart/data/item"
)

const (
	createCartQuery string = "INSERT cart SET user_id = ?, active = ?, created_date = ?"
	findCartByUserIDQuery string = "SELECT id, user_id, active, created_date FROM cart WHERE user_id = ?"
	deactivateUserCarts string = "UPDATE cart SET active = 0 WHERE id = ?"
)

// Cart Model Representing cart table in Database
type Cart struct {
	
	ID int64
	UserID string
	IsActive bool
	CreatedDate time.Time
	Items [] i.Item
}



// New Create a new Instance Of Cart for the Given UserID
func New (userID string, items [] i.Item) *Cart {
	return &Cart{UserID: userID, IsActive: true, CreatedDate: time.Now(), Items: items}
}

// Save Saves the Cart Into the Database
func (c *Cart) Save() (int64, error) {
	stmt, err := d.Db.Prepare(createCartQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(c.UserID, c.IsActive, c.CreatedDate)
	lastInserted, err := res.LastInsertId()
	for _, item := range c.Items {
		item.CartID = lastInserted
		item.Save()
	}
	markAllPreviousCartsInactive(c.UserID, lastInserted)
	return lastInserted, err
}

// FindActiveCartOfUser Finds the active cart
func FindActiveCartOfUser(userID string) *Cart {
	allCarts := FindByUserID(userID)
	for _, cart := range allCarts {
		if cart.IsActive == true {
			items := i.FindByCartID(cart.ID)
			cart.Items = items
			return &cart
		}
	}
	return nil
}

func AddItemsToCartForUser(userID string, items [] i.Item) (cartID int64) {
	activeCart := FindActiveCartOfUser(userID)
	cartID = activeCart.ID
	for _, item := range items {
		item.CartID = cartID
		item.Save()
	}
	return cartID
}

func (c *Cart) deactivateCarts() {
	stmt, err := d.Db.Prepare(deactivateUserCarts)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	stmt.Exec(c.ID)
}

func markAllPreviousCartsInactive(userID string, lastInsertedID int64) {
	previousCarts := FindByUserID(userID)
	for _, cart := range previousCarts {
		if cart.IsActive == true && cart.ID != lastInsertedID {
			cart.IsActive = false
			cart.deactivateCarts()
		}
	}
}

// FindByUserID Finds All Carts Of A User
func FindByUserID(userID string) (carts [] Cart) {
	rows, err := d.Db.Query(findCartByUserIDQuery, userID)
	if (err != nil) {
		panic(err)
	}
	for rows.Next() {
		var cart Cart
		err := rows.Scan(&cart.ID, &cart.UserID, &cart.IsActive, &cart.CreatedDate)
		if err != nil {
			panic(err)
		}
		carts = append(carts, cart)
	}
	return carts
}