package article

type Article struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
