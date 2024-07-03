package models

import "fmt"

type Item struct {
	Id      string `json:"oid"`
	Name    string `json:"name"`
	Account string `json:"account"`
	Email   string `json:"email"`
}

type CountResult struct {
	Id int64 `json:"oid"`
}

func (i *Item) String() string {
	return fmt.Sprintf("Item (id: %q, name: %q,account: %q,email: %q)", i.Id, i.Name, i.Account, i.Email)
}
