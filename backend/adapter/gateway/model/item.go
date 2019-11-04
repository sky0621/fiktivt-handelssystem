package model

type DBItem struct {
	ID           string `db:"id"`
	Name         string `db:"name"`
	Price        int    `db:"price"`
	ItemHolderID string `db:"item_holder_id"`
}

type DBItemHolder struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Nickname  string `db:"nickname"`
}
