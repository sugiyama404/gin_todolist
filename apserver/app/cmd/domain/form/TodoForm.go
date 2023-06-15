package form

type TodoRequest struct {
	ID      int    `form:"id"`
	Content string `form:"content"`
}
