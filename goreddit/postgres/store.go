package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	*ThreadStore
	*PostStore
	*CommentStore
}

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening db  %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}
	//no errors, create a new store and set err = nil
	return &Store{
		ThreadStore:  NewThreadStore(db),
		PostStore:    NewPostStore(db),
		CommentStore: NewCommentStore(db),
	}, nil // for the err
}
