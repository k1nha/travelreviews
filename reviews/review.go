package reviews

import "time"

type Review struct {
	ID          string    `json:"id"`
	Place       string    `json:"place"`
	Stars       int       `json:"stars"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
