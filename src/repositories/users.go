package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of the user repository
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

// Create inserts a new user in the database
func (r userRepository) Create(user models.User) (uint64, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

// Search returns all users that meet the passed filter
func (r userRepository) Search(userToSearch string) ([]models.User, error) {
	userToSearch = fmt.Sprintf("%%%s%%", userToSearch)

	res, err := r.db.Query(`SELECT id,
								   name,
								   nickname,
								   email,
								   createdAt 
	                          FROM users 
							 WHERE name LIKE ? 
							    OR nickname LIKE ?`, userToSearch, userToSearch)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	var users []models.User

	for res.Next() {
		var user models.User

		err := res.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetByID gets one user by the id given
func (r userRepository) GetByID(userID uint64) (models.User, error) {
	var user models.User

	res, err := r.db.Query(`SELECT id,
								   name,
								   nickname,
								   email,
								   createdAt
							  FROM users 
							 WHERE id = ?`, userID)
	if err != nil {
		return user, err
	}

	defer res.Close()

	if res.Next() {
		err = res.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			return user, err
		}
	}

	return user, err
}

// Update a user by the ID given
func (r userRepository) Update(userID uint64, user models.User) error {
	stmt, err := r.db.Prepare(`UPDATE users
								  SET name = ?,
								  	  nickname = ?,
									  email = ?
								WHERE id = ?`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Nickname, user.Email, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r userRepository) Delete(userID uint64) error {
	stmt, err := r.db.Prepare(`DELETE FROM users
								WHERE id = ?`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	return nil
}
