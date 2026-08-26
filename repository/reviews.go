package repository

import (
	"API/API/models"
)

func ListReviewsQuery(productID int) ([]models.Review, error) {

	db, err := GetData()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT id, product_id, user, rating, comment, creation_date
		FROM Reviews
		WHERE product_id = ?
	`, productID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review

	for rows.Next() {

		var review models.Review

		err := rows.Scan(
			&review.ID,
			&review.ProductID,
			&review.UserID,
			&review.Rating,
			&review.Comment,
			&review.CreationDate,
		)

		if err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}
func CreateReviewQuery(review models.Review) (models.Review, error) {

	db, err := GetData()
	if err != nil {
		return models.Review{}, err
	}
	defer db.Close()

	result, err := db.Exec(`
		INSERT INTO Reviews
		(product_id, user_id, rating, comment, creation_date)
		VALUES (?, ?, ?, ?, ?)
	`,
		review.ProductID,
		review.UserID,
		review.Rating,
		review.Comment,
		review.CreationDate,
	)

	if err != nil {
		return models.Review{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Review{}, err
	}

	review.ID = int(id)

	return review, nil
}
