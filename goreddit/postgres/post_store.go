package postgres

import (
	"fmt"

	"github.com/ivalexandru/golang_postgres_reddit001/goreddit"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jmoiron/sqlx"
)

//initializer:
func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{
		DB: db,
	}
}

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (goreddit.Post, error) {
	var p goreddit.Post
	//the first param is a refference to our Post
	//so sqlz will auto fill in the fields returned by the querry
	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf("error getting Post: %w", err)
	}
	return p, nil // nil pt eroare
}

func (s *PostStore) Posts() ([]goreddit.Post, error) {
	var pp []goreddit.Post
	if err := s.Select(&pp, `SELECT * FROM posts`); err != nil {
		return []goreddit.Post{}, fmt.Errorf("error getting Posts: %w", err)
	}
	return pp, nil
}

func (s *PostStore) CreatePost(p *goreddit.Post) error {
	if err := s.Get(p, `INSERT INTO posts VALUES ($1, $2, $3) RETURNING *`,
		p.ID,
		p.Title,
		p.Description,
	); err != nil {
		return fmt.Errorf("err creating Post: %w", err)
	}
	return nil
}

func (s *PostStore) UpdatePost(p *goreddit.Post) error {
	if err := s.Get(p, `UPDATE posts SET title = $1, description = $2 WHERE ID = $3 RETURNING *`,
		p.Title,
		p.Description,
		p.ID,
	); err != nil {
		return fmt.Errorf("err updating Post: %w", err)
	}
	return nil
}

func (s *PostStore) DeletedPost(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM posts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("err deleting Post: %w", err)
	}
	return nil
}
