package helpers

import (
	"database/sql"
	"log"
)

type Order struct {
	ID        int    `json:"id"`
	ClientID  int    `json:"client_id"`
	ManagerID int    `json:"manager_id"`
	ReportID  int    `json:"report_id"`
	Used      bool   `json:"used"`
	Cart      string `json:"cart"`
	Created   int    `json:"created"`
	Status    int    `json:"status"`
}
type CreateOrderJson struct {
	Items []Item `json:"items"`
}
type Item struct {
	PublicId  int `json:"public_id"`
	PostCount int `json:"post_count"`
}

func GetOrderById(id int64) ([]*Order, bool) {
	orders := make([]*Order, 0)
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3050)/panel")
	if err != nil {
		log.Print("Error to connect to database")
		log.Print(err)
	}
	rows, err := db.Query("SELECT * FROM orders WHERE id = ?", id)
	if err != nil {
		log.Print("Error query")
		log.Print(err)
	}
	if err != nil {
		return orders, true
	}

	defer rows.Close()
	for rows.Next() {
		order := new(Order)
		err := rows.Scan(
			&order.ID,
			&order.ClientID,
			&order.ManagerID,
			&order.ReportID,
			&order.Used,
			&order.Cart,
			&order.Created,
			&order.Status,
		)
		if err != nil {
			log.Print("Error scan row")
			log.Print(err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		log.Print("zx")
		log.Print(err)
	}

	return orders, false
}

func GetOrders() ([]*Order, bool) {
	orders := make([]*Order, 0)
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3050)/panel")
	if err != nil {
		log.Print("Error to connect to database")
		log.Print(err)
	}
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Print("Error query")
		log.Print(err)
	}
	if err != nil {
		return orders, true
	}

	defer rows.Close()
	for rows.Next() {
		order := new(Order)
		err := rows.Scan(
			&order.ID,
			&order.ClientID,
			&order.ManagerID,
			&order.ReportID,
			&order.Used,
			&order.Cart,
			&order.Created,
			&order.Status,
		)
		if err != nil {
			log.Print("Error scan row")
			log.Print(err)
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		log.Print("zx")
		log.Print(err)
	}

	return orders, false
}

func CreateOrder() (CreateOrderJson, bool) {

	// Тестовые данные
	items := make([]Item, 0)

	item1 := Item{
		PublicId:  12129,
		PostCount: 1,
	}
	item2 := Item{
		PublicId:  2222,
		PostCount: 3,
	}
	item3 := Item{
		PublicId:  333,
		PostCount: 33,
	}
	items = append(items, item1, item2, item3)

	order := CreateOrderJson{Items: items}
	return order, false
}
