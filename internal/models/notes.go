package internal

type Notes struct {
	Id     int    `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
