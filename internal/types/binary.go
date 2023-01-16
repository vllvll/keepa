package types

import "time"

type Binary struct {
	ID        int64
	Content   []byte
	Meta      string
	UpdatedAt time.Time
}
