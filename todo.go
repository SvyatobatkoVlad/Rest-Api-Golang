package Rest_Api_Golang

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"description"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}