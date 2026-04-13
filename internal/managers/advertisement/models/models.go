package advertisementModels

type Photo struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	SortOrder int    `json:"sort_order"`
	CreatedAt string `json:"created_at"`
}

// CreateAdvertisementRequest тело запроса для создания объявления
type CreateAdvertisementRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	PhotoFiles  [][]byte
}

// AdvertisementResponse тело ответа с данными объявления
type AdvertisementResponse struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
	Quantity    int     `json:"quantity"`
	Photos      []Photo `json:"photos"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
