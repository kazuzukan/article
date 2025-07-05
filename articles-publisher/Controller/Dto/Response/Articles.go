package Response

type ArticlesData struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	AuthorName string `json:"author_name"`
	CreatedAt  string `json:"created_at"`
}

type Articles struct {
	TotalData    int            `json:"total_data"`
	ArticlesList []ArticlesData `json:"articles_list"`
}
