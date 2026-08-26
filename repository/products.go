package repository

import (
	"API/API/models"
)

func GetProductsQuery() ([]models.Product, error) {
	db, err := GetData()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT id, name, description, price, category, stock
		FROM Products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Category,
			&product.Stock,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
func CreateProductQuery(product models.Product) (int, error) {
	db, err := GetData()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	result, err := db.Exec(`
		INSERT INTO Products
		(name, description, price, category, stock)
		VALUES (?, ?, ?, ?, ?)
	`,
		product.Name,
		product.Description,
		product.Price,
		product.Category,
		product.Stock,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
func UpdateProductQuery(id int, update models.ProductUpdateReq) (models.Product, error) {

	db, err := GetData()
	if err != nil {
		return models.Product{}, err
	}
	defer db.Close()

	_, err = db.Exec(`
		UPDATE Products
		SET price = ?, stock = ?
		WHERE id = ?
	`,
		update.Price,
		update.Stock,
		id,
	)

	if err != nil {
		return models.Product{}, err
	}

	var product models.Product

	err = db.QueryRow(`
		SELECT id, name, description, price, category, stock
		FROM Products
		WHERE id = ?
	`, id).Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Category,
		&product.Stock,
	)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
