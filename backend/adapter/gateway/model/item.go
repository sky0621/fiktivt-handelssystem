package model

type DBItem struct {
	ID           string `db:"id"`
	Name         string `db:"name"`
	Price        int    `db:"price"`
	ItemHolderID string `db:"item_holder_id"`
}
