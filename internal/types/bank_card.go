package types

import "time"

type BankCard struct {
	ID        int64
	Number    string
	Holder    string
	CVV       string
	Meta      string
	UpdatedAt time.Time
}
