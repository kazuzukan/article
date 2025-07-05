package Articles

import (
	"articles-consumer/Controller/Dto/Request"
	"articles-consumer/Controller/Dto/Response"
	"database/sql"
	"fmt"
)

func (a article) CreateArticle(dataRequest Request.CreateArticle) (err error) {
	tx, err := a.dbCon.PostgreSQLConnection().Begin()
	if err != nil {
		return
	}

	authorId, err := a.insertAndGetAuthor(tx, dataRequest.AuthorName)
	if err != nil {
		tx.Rollback()
		return
	}

	query := `INSERT INTO articles.t_articles (author_id, title, body) VALUES ($1, $2, $3)`
	_, err = tx.Exec(query, authorId, dataRequest.Title, dataRequest.Body)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return
}

func (a article) insertAndGetAuthor(tx *sql.Tx, authorName string) (authorId int, err error) {
	queryGet := `SELECT id FROM users.t_author WHERE UPPER(name) = UPPER($1)`
	err = tx.QueryRow(queryGet, authorName).Scan(&authorId)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if err == sql.ErrNoRows {
		queryInsert := `INSERT INTO users.t_author (name) VALUES ($1) RETURNING id`
		err = tx.QueryRow(queryInsert, authorName).Scan(&authorId)
		if err != nil {
			return
		}
	}

	return
}

func (a article) GetArticles(params Request.GetArticles) (articles Response.Articles, err error) {
	connection := a.dbCon.PostgreSQLConnection()
	limit := 5
	offset := (params.Page - 1) * limit
	args := []interface{}{params.Keyword, limit, offset}
	var filterByAuthor string
	if params.AuthorName != "" {
		filterByAuthor = fmt.Sprintf(" AND au.name = $%d", len(args)+1)
		args = append(args, params.AuthorName)
	}

	query := `SELECT au.name, ta.title, ta.body, TO_CHAR(ta.created_at, 'DD-MM-YYYY HH24:MI'), COUNT(1) OVER()
				FROM articles.t_articles ta
				INNER JOIN users.t_author au ON au.id = ta.author_id
				WHERE (
					UPPER(ta.title) LIKE UPPER('%' || $1 || '%') OR 
					UPPER(ta.body) LIKE UPPER('%' || $1 || '%')
				) ` + filterByAuthor + `
				ORDER BY ta.created_at DESC
				LIMIT $2 OFFSET $3`

	rows, err := connection.Query(query, args...)
	if err != nil {
		return
	}

	for rows.Next() {
		var articlesData Response.ArticlesData
		err = rows.Scan(&articlesData.AuthorName, &articlesData.Title, &articlesData.Body, &articlesData.CreatedAt, &articles.TotalData)
		if err != nil {
			return
		}

		articles.ArticlesList = append(articles.ArticlesList, articlesData)
	}

	if len(articles.ArticlesList) == 0 {
		articles.ArticlesList = []Response.ArticlesData{}
	}

	return
}
