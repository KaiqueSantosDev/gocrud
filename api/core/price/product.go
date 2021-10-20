package price

type Product struct {
	ID          int64   `gorm:"primary_key" json:"id"`
	Name        string  `gorm:"size:100;not null;" json:"name"`
	Description string  `gorm:"size:150;not null;"  json:"description"`
	Price       float32 `gorm:"not null;" json:"price"`
}
