package repositories

import (
	"database/sql"
	"github.com/vllvll/keepa/internal/types"
	"time"
)

type BankCard struct {
	db *sql.DB
}

type BankCardInterface interface {
	Get(id int64) (card types.BankCard, err error)
	Create(bankCard types.BankCard, userId int64) (int64, error)
	Update(bankCard types.BankCard) error
	Delete(bankCardID int64) error
}

func NewBankCardRepository(db *sql.DB) BankCardInterface {
	return &BankCard{db: db}
}

func (b *BankCard) Get(id int64) (bankCard types.BankCard, err error) {
	err = b.db.QueryRow("SELECT id, number, holder, cvv, meta, updated_at FROM bank_cards WHERE id = $1 LIMIT 1", id).
		Scan(&bankCard.ID, &bankCard.Number, &bankCard.Holder, &bankCard.CVV, &bankCard.Meta, &bankCard.UpdatedAt)
	if err != nil {
		return types.BankCard{}, err
	}

	return bankCard, nil
}

func (b *BankCard) Create(bankCard types.BankCard, userId int64) (int64, error) {
	var bankCardID int64

	row := b.db.QueryRow(
		"INSERT INTO bank_cards (number, holder, cvv, meta, user_id, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		bankCard.Number,
		bankCard.Holder,
		bankCard.CVV,
		bankCard.Meta,
		userId,
		time.Now(),
	)

	err := row.Scan(&bankCardID)
	if err != nil {
		return bankCardID, err
	}

	return bankCardID, nil
}

func (b *BankCard) Update(bankCard types.BankCard) error {
	_, err := b.db.Exec(
		"UPDATE bank_cards SET number = $2, holder = $3, cvv = $4, meta = $5, updated_at = $6 WHERE id = $1",
		bankCard.ID,
		bankCard.Number,
		bankCard.Holder,
		bankCard.CVV,
		bankCard.Meta,
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *BankCard) Delete(bankCardID int64) error {
	_, err := b.db.Exec(
		"DELETE FROM bank_cards WHERE id = $1",
		bankCardID,
	)
	if err != nil {
		return err
	}

	return nil
}
