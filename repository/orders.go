package repository

import (
	"API/API/models"

	"fmt"
)

func CreateOrderQuery(order models.Order) (models.Order, error) {

	db, err := GetData()
	if err != nil {
		return models.Order{}, err
	}
	defer db.Close()

	var totalPrice float64

	for _, item := range order.Items {

		var price float64
		var stock int

		err = db.QueryRow(`
			SELECT price, stock
			FROM Products
			WHERE id = ?
		`, item.ProductID).Scan(&price, &stock)

		if err != nil {
			return models.Order{}, fmt.Errorf("product not found")
		}

		if item.Quantity <= 0 || item.Quantity > stock {
			return models.Order{}, fmt.Errorf("invalid quantity")
		}

		totalPrice += price * float64(item.Quantity)
	}

	order.Total_Price = totalPrice

	result, err := db.Exec(`
		INSERT INTO Orders
		(user_id, total_price, status)
		VALUES (?, ?, ?)
	`,
		order.UserID,
		order.Total_Price,
		order.Status,
	)

	if err != nil {
		return models.Order{}, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return models.Order{}, err
	}

	order.ID = int(orderID)

	for _, item := range order.Items {

		_, err = db.Exec(`
			INSERT INTO OrderItems
			(order_id, product_id, quantity)
			VALUES (?, ?, ?)
		`,
			order.ID,
			item.ProductID,
			item.Quantity,
		)

		if err != nil {
			return models.Order{}, err
		}
	}

	return order, nil
}
func DeleteOrderQuery(userID int, orderID int) (models.Order, error) {

	db, err := GetData()
	if err != nil {
		return models.Order{}, err
	}
	defer db.Close()

	result, err := db.Exec(`
		UPDATE Orders
		SET status = ?
		WHERE id = ? AND user_id = ?
	`, "cancelled", orderID, userID)

	if err != nil {
		return models.Order{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Order{}, err
	}

	if rowsAffected == 0 {
		return models.Order{}, fmt.Errorf("order not found")
	}

	var order models.Order

	err = db.QueryRow(`
		SELECT id, user_id, total_price, status
		FROM Orders
		WHERE id = ? AND user_id = ?
	`, orderID, userID).Scan(
		&order.ID,
		&order.UserID,
		&order.Total_Price,
		&order.Status,
	)

	if err != nil {
		return models.Order{}, err
	}

	rows, err := db.Query(`
		SELECT product_id, quantity
		FROM OrderItems
		WHERE order_id = ?
	`, orderID)

	if err != nil {
		return models.Order{}, err
	}
	defer rows.Close()

	for rows.Next() {

		var item models.OrderItem

		err = rows.Scan(
			&item.ProductID,
			&item.Quantity,
		)

		if err != nil {
			return models.Order{}, err
		}

		order.Items = append(order.Items, item)
	}
	if err = rows.Err(); err != nil {
		return models.Order{}, err
	}

	return order, nil
}
