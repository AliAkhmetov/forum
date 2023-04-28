package repository

import (
	"database/sql"
	"fmt"

	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

type likeSQL struct {
	db *sql.DB
}

// New create new database.
func NewlikeSQL(db *sql.DB) *likeSQL {
	return &likeSQL{db: db}
}

func (r *likeSQL) CreateLike(like models.Like) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (author_id, post_id, comment_id, type) values (?,?,?,?) RETURNING id`, likesTable)

	row := r.db.QueryRow(query, like.AuthorID, like.PostID, like.CommentID, like.Type)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *likeSQL) GetLikesByPostId(postId int) ([]models.Like, error) {
	var allLikes []models.Like

	query := fmt.Sprintf("SELECT id, author_id, post_id, comment_id, type FROM %s WHERE post_id = ?", likesTable)
	err := r.db.QueryRow(query, postId).Scan(&allLikes)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can't get all likes of this post: %w", err)
	}
	return allLikes, err
}

func (r *likeSQL) GetLikesByCommentId(commentId int) ([]models.Like, error) {
	var allLikes []models.Like

	query := fmt.Sprintf("SELECT id, author_id, post_id, comment_id, type FROM %s WHERE comment_id = ?", likesTable)
	err := r.db.QueryRow(query, commentId).Scan(&allLikes)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("can't get all likes of this comment: %w", err)
	}
	return allLikes, err
}
