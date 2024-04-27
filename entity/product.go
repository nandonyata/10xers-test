package entity

type Product struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Type   string `json:"type"`
	Price  int    `json:"price"`
	Stock  int    `json:"stock"`
}
