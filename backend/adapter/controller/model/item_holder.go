package model

import "fmt"

type ItemHolder struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Nickname *string `json:"nickname"`
}

func (i *ItemHolder) String() string {
	return fmt.Sprintf("[custom] %#v\n", i)
}
