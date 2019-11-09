package model

type Item struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	ItemHolderID string `json:"itemHolderID"`
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

type ItemHolderConnection struct {
	TotalCount int              `json:"totalCount"`
	Edges      []ItemHolderEdge `json:"edges"`
	PageInfo   *PageInfo        `json:"pageInfo"`
}

type ItemHolderEdge struct {
	Cursor string      `json:"cursor"`
	Node   *ItemHolder `json:"node"`
}

type SearchItemHolderCondition struct {
	Nickname *string `json:"nickname"`
}
