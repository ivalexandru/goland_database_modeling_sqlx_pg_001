package goreddit

import uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

// go get github.com/google/uuid

//thread = collection of posts (reddit)
type Thread struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

type Post struct {
	ID       uuid.UUID `db:"id"`
	ThreadID uuid.UUID `db:"thread_id"`
	Title    string    `db:"title"`
	Content  string    `db:"content"`
	Votes    int       `db:"votes"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	PostID  uuid.UUID `db:"post_id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
}

type ThreadStore interface {
	//returns a Thread and an error
	Thread(id uuid.UUID) (Thread, error)
	Threads() ([]Thread, error) //all threads
	CreateThread(t *Thread) error
	UpdateThread(t *Thread) error
	DeleteThread(id uuid.UUID) error
}

type PostStore interface {
	//returns a Thread and an error
	Post(id uuid.UUID) (Post, error)
	PostsByThread(threadID uuid.UUID) ([]Post, error) //all threads
	CreatePost(t *Post) error
	UpdatePost(t *Post) error
	DeletePost(id uuid.UUID) error
}

type CommentStore interface {
	//returns a Thread and an error
	Comment(id uuid.UUID) (Comment, error)
	CommentsByPost(postID uuid.UUID) ([]Comment, error) //all threads
	CreateComment(t *Comment) error
	UpdateComment(t *Comment) error
	DeleteComment(id uuid.UUID) error
}

//will help for dependency injection
type Store interface {
	ThreadStore
	PostStore
	CommentStore
}

// https://www.youtube.com/watch?v=-aRJfl44mQc&t=675s
