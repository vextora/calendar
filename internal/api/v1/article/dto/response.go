package dto

type ArticleResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Content string `json:"content,omitempty"`
}
