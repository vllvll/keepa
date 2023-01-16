package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"

	"github.com/vllvll/keepa/internal/types"
)

type Text struct {
	db *sql.DB
}

type TextInterface interface {
	Get(id int64) (text types.Text, err error)
	Create(text types.Text, userId int64) (int64, error)
	Update(text types.Text) error
	Delete(textID int64) error
}

func NewTextRepository(db *sql.DB) TextInterface {
	return &Text{db: db}
}

func (b *Text) Get(id int64) (text types.Text, err error) {
	err = b.db.QueryRow("SELECT id, content, meta, updated_at FROM texts WHERE id = $1 LIMIT 1", id).
		Scan(&text.ID, &text.Content, &text.Meta, &text.UpdatedAt)

	if err, ok := err.(*pq.Error); ok {
		if pgerrcode.IsConnectionException(string(err.Code)) {
			log.Fatalf("Error with database: %v", err)
		}

		return types.Text{}, err
	}

	return text, nil
}

func (b *Text) Create(text types.Text, userId int64) (int64, error) {
	var textID int64

	row := b.db.QueryRow(
		"INSERT INTO texts (content, meta, user_id, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		text.Content,
		text.Meta,
		userId,
		time.Now(),
	)

	err := row.Scan(&textID)
	if err, ok := err.(*pq.Error); ok {
		if pgerrcode.IsConnectionException(string(err.Code)) {
			log.Fatalf("Error with database: %v", err)
		}

		return textID, err
	}

	return textID, nil
}

func (b *Text) Update(text types.Text) error {
	_, err := b.db.Exec(
		"UPDATE texts SET content = $2, meta = $3, updated_at = $4 WHERE id = $1",
		text.ID,
		text.Content,
		text.Meta,
		time.Now(),
	)
	if err, ok := err.(*pq.Error); ok {
		if pgerrcode.IsConnectionException(string(err.Code)) {
			log.Fatalf("Error with database: %v", err)
		}

		return err
	}

	return nil
}

func (b *Text) Delete(textID int64) error {
	_, err := b.db.Exec(
		"DELETE FROM texts WHERE id = $1",
		textID,
	)

	if err, ok := err.(*pq.Error); ok {
		if pgerrcode.IsConnectionException(string(err.Code)) {
			log.Fatalf("Error with database: %v", err)
		}

		return err
	}

	return nil
}
