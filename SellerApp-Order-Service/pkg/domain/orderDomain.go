package domain

type Order struct {
	ID           string  `json:"id" gorm:"primaryKey;unique"`
	Status       string  `json:"status"`
	Item_id      string  `json:"items" `
	Total        float64 `json:"total"`
	CurrencyUnit string  `json:"currencyUnit"`
}

type Item struct {
	ID          string  `json:"id" gorm:"primaryKey;unique"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
