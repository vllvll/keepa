package types

import "time"

type Text struct {
	ID        int64
	Content   string
	Meta      string
	UpdatedAt time.Time
}
