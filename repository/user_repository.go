package repository

import (
	"database/sql"
	"wheels-api/model"
)

type UserRepository interface {
	CreateUser(user model.User) (int64, error)
	GetUserByEmail(email string) (*model.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}

func (ur *userRepository) CreateUser(user model.User) (int64, error) {
	var id int64
	err := ur.db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *userRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
