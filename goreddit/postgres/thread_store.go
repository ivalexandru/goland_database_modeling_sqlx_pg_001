package postgres

import (
	"fmt"

	"github.com/ivalexandru/golang_postgres_reddit001/goreddit"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jmoiron/sqlx"
)

//embedding the type in our struct is done by omitting the fieldname:
//this means our ThreadStore type borrows all of the functionality
// and methods of the sqlx.DB type
//so we call methods right on our ThreadStore
type ThreadStore struct {
	*sqlx.DB
}

func NewThreadStore(db *sqlx.DB) *ThreadStore {
	return &ThreadStore{
		DB: db,
	}
}

// ctrl+shift+p in vscode
// search for
//generate interface stubs
// apoi tastezi in cazuta respectiva
// s*ThreadStore the interface we're implementing, aka:
// s*ThreadStore goreddit.ThreadStore

func (s *ThreadStore) Thread(id uuid.UUID) (goreddit.Thread, error) {
	var t goreddit.Thread
	//the first param is a refference to our thread
	//so sqlz will auto fill in the fields returned by the querry
	if err := s.Get(&t, `SELECT * FROM threads WHERE id = $1`, id); err != nil {
		return goreddit.Thread{}, fmt.Errorf("error getting thread: %w", err)
	}
	return t, nil // nil pt eroare
}

func (s *ThreadStore) Threads() ([]goreddit.Thread, error) {
	var tt []goreddit.Thread
	if err := s.Select(&tt, `SELECT * FROM threads`); err != nil {
		return []goreddit.Thread{}, fmt.Errorf("error getting threads: %w", err)
	}
	return tt, nil
}

func (s *ThreadStore) CreateThread(t *goreddit.Thread) error {
	if err := s.Get(t, `INSERT INTO threads VALUES ($1, $2, $3) RETURNING *`,
		t.ID,
		t.Title,
		t.Description,
	); err != nil {
		return fmt.Errorf("err creating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) UpdateThread(t *goreddit.Thread) error {
	if err := s.Get(t, `UPDATE threads SET title = $1, description = $2 WHERE ID = $3 RETURNING *`,
		t.Title,
		t.Description,
		t.ID,
	); err != nil {
		return fmt.Errorf("err updating thread: %w", err)
	}
	return nil
}

func (s *ThreadStore) DeletedThread(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM threads WHERE id = $1`, id); err != nil {
		return fmt.Errorf("err deleting thread: %w", err)
	}
	return nil
}
