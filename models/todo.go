package models

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	ID    int   `json:"id"`
	Todos []int `json:"todos"`
}
