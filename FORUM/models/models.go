package models

type User struct {
	Id       int    `json:"id" db:"id"`
	UserName string `json:"userName"  db:"username"`
	Email    string `json:"email"  db:"email"`
	PassHash string `json:"email"  db:"password_hash"`
}

type Post struct {
	Id         int    `json:"id"  db:"id"`
	AuthorID   string `json:"authorId"  db:"id"`
	Text       string `json:"text"  db:"id"`
	LikesCount int    `json:"likesCount" db:"id"`
}

type Comment struct {
	Id         int    `json:"id"  db:"id"`
	AuthorID   string `json:"authorId"  db:"id"`
	PostID     string `json:"postID"  db:"id"`
	Text       string `json:"text"  db:"id"`
	LikesCount int    `json:"likesCount"  db:"id"`
}

type Like struct {
	Id        int    `json:"id"  db:"id"`
	AuthorID  string `json:"authorId"  db:"author_id"`
	PostID    string `json:"postID"  db:"post_id"`
	CommentID string `json:"commentID"  db:"comment_id"`
	Type      int    `json:"type"  db:"type"`
}
