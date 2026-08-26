package repository

import (
	"API/API/models"
)

func CreateUserQuery(user models.User) (int, error) {

	db, err := GetData()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(`
		INSERT INTO Users
		(name, email, password, role)
		VALUES (?, ?, ?, ?)
	`,
		user.Name,
		user.Email,
		hashedPassword,
		user.Role,
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
