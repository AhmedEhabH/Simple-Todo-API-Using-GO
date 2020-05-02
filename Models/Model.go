package Models

// Todo Table
type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TableName return todo
func (b *Todo) TableName() string {
	return "todo"
}
