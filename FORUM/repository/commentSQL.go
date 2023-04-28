package repository

import (
	"database/sql"
	"fmt"

	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

type commentSQL struct {
	db *sql.DB
}

// New create new database.
func NewCommentSQL(db *sql.DB) *commentSQL {
	return &commentSQL{db: db}
}

func (r *commentSQL) CreateComment(comment models.Comment) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (author_id, post_id, text, likes-count) values (?,?,?,?) RETURNING id`, commentsTable)

	row := r.db.QueryRow(query, comment.AuthorID, comment.PostID, comment.Text, comment.LikesCount)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *commentSQL) GetCommentByPostId(postId int) ([]models.Comment, error) {
	var allComments []models.Comment

	query := fmt.Sprintf("SELECT id, author_id, post_id, text,likes_count FROM %s WHERE post_id = ?", commentsTable)
	err := r.db.QueryRow(query, postId).Scan(&allComments)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can't get all Comments: %w", err)
	}
	return allComments, err
}
