package repository

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"forum/models"

	_ "github.com/mattn/go-sqlite3"
)

const (
	usersTable             = "users"
	postsTable             = "posts"
	categoriesTable        = "categories"
	categoriesToPostsTable = "categories_posts"
	commentsTable          = "comments"
	likesTable             = "likes"
)

type Storage struct {
	Db *sql.DB
}

type Authorization interface {
	CreateUser(User models.User) (int, error)
	GetUser(UserName, PassHash string) (models.User, error)
}

type Posts interface {
	CreatePost(post models.Post) (int, error)
	GetAll() ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (int, error)
	GetCommentByPostId(postId int) ([]models.Comment, error)
}
type Likes interface {
	CreateLike(like models.Like) (int, error)
	GetLikesByPostId(postId int) ([]models.Like, error)
	GetLikesByCommentId(commentId int) ([]models.Like, error)
}

type Repository struct {
	Authorization
	Posts
	Comments
	Likes
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("Can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Can't connect to database: %w", err)
	}

	return &Storage{Db: db}, nil
}

// Init all
func (s *Storage) Init(initSqlFileName string) error {
	file, err := ioutil.ReadFile(initSqlFileName)
	if err != nil {
		log.Fatalf("Can't read SQL file %v", err)
	}

	// Execute all
	_, err = s.Db.Exec(string(file))
	if err != nil {
		log.Fatalf("DB init error: %v", err)
	}
	return nil
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
		Posts:         NewPostsSQL(db),
		Comments:      NewCommentSQL(db),
		Likes:         NewlikeSQL(db),
	}
}
