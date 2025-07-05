package Request

type CreateArticles struct {
	AuthorName string `json:"author_name"`
	Title      string `json:"title"`
	Body       string `json:"body"`
}

type GetArticles struct {
	Page       int    `json:"page"`
	AuthorName string `json:"author_name"`
	Keyword    string `json:"keyword"`
}
