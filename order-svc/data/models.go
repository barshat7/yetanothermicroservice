package data

import (
	"time"
	"database/sql"
	_ "fmt"
)

type Order struct {

	ID int64
	UserID string
	CreatedDate time.Time
	TotalPrice float64
}

func SaveOrder(userID string, totalPrice float64) (int64, error) {

	sql := "INSERT orders SET user_id = ?, total_price = ?, created_date = ?"
	stmt, err := Db.Prepare(sql)
	if (err != nil) {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(userID, totalPrice, time.Now())
	lastInserted, err := res.LastInsertId()
	return lastInserted, err
}

func FindOrderByID(orderID int64) *Order {
	sql := "SELECT id, user_id, created_date, total_price FROM orders WHERE id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Query(orderID)
	defer closeRows(res)
	if err != nil {
		panic(err)
	}
	var order Order
	if res.Next() {
		res.Scan(&order.ID, &order.UserID, &order.CreatedDate, &order.TotalPrice)
	}
	return &order
}

func FindOrdersByUserID(userID string) (orders [] Order) {
	sql := "SELECT id, user_id, created_date, total_price FROM orders WHERE user_id = ?"
	rows, err := Db.Query(sql, userID)
	if (err != nil) {
		panic(err)
	}
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.UserID, &order.CreatedDate, &order.TotalPrice)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	return orders
}

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}