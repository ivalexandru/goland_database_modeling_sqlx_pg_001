package postgres

import (
	"fmt"

	"github.com/ivalexandru/golang_postgres_reddit001/goreddit"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jmoiron/sqlx"
)

func NewCommentStore(db *sqlx.DB) *CommentStore {
	return &CommentStore{
		DB: db,
	}
}

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id uuid.UUID) (goreddit.Comment, error) {
	var c goreddit.Comment
	//the first param is a refference to our Comment
	//so sqlz will auto fill in the fields returned by the querry
	if err := s.Get(&c, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return goreddit.Comment{}, fmt.Errorf("error getting Comment: %w", err)
	}
	return c, nil // nil pt eroare
}

func (s *CommentStore) Comments() ([]goreddit.Comment, error) {
	var tt []goreddit.Comment
	if err := s.Select(&tt, `SELECT * FROM comments`); err != nil {
		return []goreddit.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}
	return tt, nil
}

func (s *CommentStore) CreateComment(c *goreddit.Comment) error {
	if err := s.Get(c, `INSERT INTO comments VALUES ($1, $2, $3) RETURNING *`,
		c.ID,
		c.Title,
		c.Description,
	); err != nil {
		return fmt.Errorf("err creating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) UpdateComment(c *goreddit.Comment) error {
	if err := s.Get(c, `UPDATE comments SET title = $1, description = $2 WHERE ID = $3 RETURNING *`,
		c.Title,
		c.Description,
		c.ID,
	); err != nil {
		return fmt.Errorf("err updating comment: %w", err)
	}
	return nil
}

func (s *CommentStore) DeletedComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1`, id); err != nil {
		return fmt.Errorf("err deleting comment: %w", err)
	}
	return nil
}
