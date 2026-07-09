package articles

import "time"

type Article struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Journalist  string    `json:"journalist"`
	CreatedAt   time.Time `json:"createdAt"`
	Corruption  float64   `json:"corruption"`
	Credibility float64   `json:"credibility"`
}
