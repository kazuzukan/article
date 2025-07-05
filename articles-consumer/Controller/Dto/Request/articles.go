package Request

type CreateArticle struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	AuthorName string `json:"author_name"`
	AuthorId   int    `json:"author_id"`
}

type GetArticles struct {
	Page       int    `json:"page"`
	AuthorName string `json:"author_name"`
	Keyword    string `json:"keyword"`
}
