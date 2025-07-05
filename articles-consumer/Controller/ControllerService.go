package Controller

import (
	"articles-consumer/Controller/Articles"
	"articles-consumer/Controller/Dto"
)

// consumer
type ControllerConsumer struct {
	Articles Articles.ArticlesControllerInterface
}

func InitController(u Dto.Utilities) ControllerConsumer {
	return ControllerConsumer{
		Articles: Articles.NewController(u),
	}
}

// API
type ControllerAPI struct {
	Articles Articles.ControllerAPI
}

func InitControllerAPI(u Dto.Utilities) ControllerAPI {
	return ControllerAPI{
		Articles: Articles.NewControllerAPI(u),
	}
}
