package model

import "fmt"

type DBItem struct {
	ID           string `db:"id"`
	Name         string `db:"name"`
	Price        int    `db:"price"`
	ItemHolderID string `db:"item_holder_id"`
}

func (i *DBItem) String() string {
	return fmt.Sprintf("[adapter/gateway/model/DBItem]ID:%s, Name:%s, Price:%d, ItemHolderID:%s", i.ID, i.Name, i.Price, i.ItemHolderID)
}

type DBItemHolder struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Nickname  string `db:"nickname"`
}

func (i *DBItemHolder) String() string {
	return fmt.Sprintf("[adapter/gateway/model/DBItemHolder]ID:%s, FirstName:%s, LastName:%s, Nickname:%s", i.ID, i.FirstName, i.LastName, i.Nickname)
}
