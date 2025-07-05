package Service

import (
	"articles-publisher/Config"
	"articles-publisher/Controller"
	"articles-publisher/Controller/Dto"
	"articles-publisher/Routes"
	"articles-publisher/Service/Articles"
	"articles-publisher/Service/RabbitMQ"
	"github.com/go-chi/chi"
	"net/http"
)

func GenerateUtilities() (utilities Dto.Utilities) {
	connection := Config.BuildConnection()
	rabbitConfigPublish := RabbitMQ.GetRabbitMQService(connection.RabbitMQConnection())
	utilities = Dto.Utilities{
		RabbitMQ:        rabbitConfigPublish,
		ArticlesService: Articles.NewArticleService(),
	}

	return
}

func InitApplication() {
	utilities := GenerateUtilities()
	router := Routes.Routes{
		Chi:        chi.NewRouter(),
		Controller: Controller.InitController(utilities),
	}

	err := http.ListenAndServe(":9090", router.GetRoutes())
	if err != nil {
		panic(err)
	}
}
