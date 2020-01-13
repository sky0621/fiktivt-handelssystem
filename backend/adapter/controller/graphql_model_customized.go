package controller

type Item struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	ItemHolderID string `json:"itemHolderID"`
}

func (Item) IsNode() {}

type ItemHolder struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Nickname  *string `json:"nickname"`
	//HoldItems []Item  `json:"holdItems"`
}

func (ItemHolder) IsNode() {}

func (h *ItemHolder) Name() *string {
	n := h.FirstName + " " + h.LastName
	return &n
}
