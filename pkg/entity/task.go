package entity

type TaskList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TaskItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
