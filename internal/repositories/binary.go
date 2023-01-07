package repositories

import (
	"database/sql"
	"github.com/vllvll/keepa/internal/types"
	"time"
)

type Binary struct {
	db *sql.DB
}

type BinaryInterface interface {
	Get(id int64) (text types.Binary, err error)
	Create(text types.Binary, userId int64) (int64, error)
	Update(text types.Binary) error
	Delete(textID int64) error
}

func NewBinaryRepository(db *sql.DB) BinaryInterface {
	return &Binary{db: db}
}

func (b *Binary) Get(id int64) (text types.Binary, err error) {
	err = b.db.QueryRow("SELECT id, content, meta, updated_at FROM binaries WHERE id = $1 LIMIT 1", id).
		Scan(&text.ID, &text.Content, &text.Meta, &text.UpdatedAt)
	if err != nil {
		return types.Binary{}, err
	}

	return text, nil
}

func (b *Binary) Create(text types.Binary, userId int64) (int64, error) {
	var textID int64

	row := b.db.QueryRow(
		"INSERT INTO binaries (content, meta, user_id, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		text.Content,
		text.Meta,
		userId,
		time.Now(),
	)

	err := row.Scan(&textID)
	if err != nil {
		return textID, err
	}

	return textID, nil
}

func (b *Binary) Update(text types.Binary) error {
	_, err := b.db.Exec(
		"UPDATE binaries SET content = $2, meta = $3, updated_at = $4 WHERE id = $1",
		text.ID,
		text.Content,
		text.Meta,
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *Binary) Delete(textID int64) error {
	_, err := b.db.Exec(
		"DELETE FROM binaries WHERE id = $1",
		textID,
	)
	if err != nil {
		return err
	}

	return nil
}
