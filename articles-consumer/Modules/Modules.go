package Modules

import (
	"articles-consumer/Modules/Articles"
	"articles-consumer/Repository"
)

type Module struct {
	Articles Articles.ArticleModule
}

func InitModule(repo Repository.Repository) Module {
	return Module{
		Articles: Articles.NewModule(repo),
	}
}
