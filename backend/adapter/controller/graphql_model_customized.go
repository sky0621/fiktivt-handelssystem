package controller

import (
	"encoding/base64"
	"fmt"
)

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

func (h *ItemHolder) GetCursor(key string) *string {
	switch key {
	case "id":
		return EncodeCursor("id", h.ID)
	case "firstName":
		return EncodeCursor("firstName", h.FirstName)
	case "lastName":
		return EncodeCursor("lastName", h.LastName)
	case "nickname":
		return EncodeCursor("nickname", h.Nickname)
	}
	return nil
}

func EncodeCursor(key string, val interface{}) *string {
	if key == "" {
		return nil
	}
	if val == nil {
		return nil
	}
	cursor := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s=%v", key, val)))
	return &cursor
}
