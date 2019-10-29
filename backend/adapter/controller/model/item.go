package model

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type ItemHolder struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Nickname  *string `json:"nickname"`
}

func (h *ItemHolder) Name() *string {
	n := h.FirstName + " " + h.LastName
	return &n
}
