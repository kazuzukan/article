package Controller

import (
	"articles-publisher/Controller/Articles"
	"articles-publisher/Controller/Dto"
)

type Controller struct {
	Articles Articles.ArticlesControllerInterface
}

func InitController(u Dto.Utilities) Controller {
	return Controller{
		Articles: Articles.NewController(u),
	}
}
