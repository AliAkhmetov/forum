package repository

import (
	"database/sql"
	"fmt"

	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

type postSQL struct {
	db *sql.DB
}

// New create new database.
func NewPostsSQL(db *sql.DB) *postSQL {
	return &postSQL{db: db}
}

func (r *postSQL) CreatePost(post models.Post) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (author-id, text, likes-count) values (?,?,?) RETURNING id`, postsTable)

	row := r.db.QueryRow(query, post.AuthorID, post.Text, post.LikesCount)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *postSQL) GetAll() ([]models.Post, error) {
	var allPosts []models.Post

	query := fmt.Sprintf("SELECT id author-id text	likes-count FROM %s", postsTable)
	err := r.db.QueryRow(query).Scan(&allPosts)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can't get all posts: %w", err)
	}
	return allPosts, err
}
