package repository

import (
	"database/sql"
	"fmt"

	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

type usersSQL struct {
	db *sql.DB
}

// New create new database.
func NewAuthSQL(db *sql.DB) *usersSQL {
	return &usersSQL{db: db}
}

func (r *usersSQL) CreateUser(User models.User) (int, error) {
	fmt.Println("User in SQL", User)

	var id int
	query := fmt.Sprintf(`INSERT INTO %s ( email, username, password_hash) values (?,?,?) RETURNING id`, usersTable)
	fmt.Println(query)

	row := r.db.QueryRow(query, User.Email, User.UserName, User.PassHash)
	if err := row.Scan(&id); err != nil {
		fmt.Println("EEERRRRORRR,", err.Error())

		return 0, err
	}
	fmt.Println("!!!!! ID", id)

	return id, nil
}

func (r *usersSQL) GetUser(Email, PassHash string) (models.User, error) {
	fmt.Println("GET User SQL")
	fmt.Println(Email, PassHash)
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=? AND password_hash=?", usersTable)
	err := r.db.QueryRow(query, Email, PassHash).Scan(&user.Id, &user.Email, &user.UserName, &user.PassHash)
	if err == sql.ErrNoRows {
		return user, err
	}
	if err != nil {
		return user, fmt.Errorf("can't get user: %w", err)
	}
	return user, nil
}
