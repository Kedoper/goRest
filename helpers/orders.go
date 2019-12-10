package helpers

type Order struct {
	ID int `json:"id"`
}

func GetOrderById(id int64) int {
	return int(id)
}
