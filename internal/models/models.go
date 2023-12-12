// Package models contains data structures common through the program.
package models

type Post struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type Data struct {
	Users []Post `json:"users"`
}
